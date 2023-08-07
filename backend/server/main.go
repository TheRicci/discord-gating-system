package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mattn/go-colorable"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

type userInfo struct {
	userID      string
	guildID     string
	interaction *discordgo.Interaction
	*time.Time
}
type awaitingConnectionMessageToUserMap map[string]*userInfo
type awaitingConnectionMessageToUser struct {
	m awaitingConnectionMessageToUserMap
	*sync.RWMutex
}

type userToAddressMap map[string][]string
type userToAddress struct {
	m userToAddressMap
	*sync.RWMutex
}

type contract struct {
	Address string
	tokenType
}

type roleToContractsMap map[string][]contract
type roleToContracts struct {
	m roleToContractsMap
	*sync.RWMutex
}

type guildIDtoRolesMap map[string][]string
type guildIDtoRoles struct {
	m guildIDtoRolesMap
	*sync.RWMutex
}

type contractsMap map[string]contractDB
type contracts struct {
	m contractsMap
	*sync.RWMutex
}

type guildsMap map[string]guild
type guilds struct {
	m guildsMap
	*sync.RWMutex
}

type Bot struct {
	*discordgo.Session
	*gin.Engine

	*guilds
	*contracts
	*awaitingConnectionMessageToUser
	*userToAddress
	*roleToContracts
	*guildIDtoRoles
	*database
}

type connectBind struct {
	Message       string
	SignedMessage string `json:"signed_message"`
	Wallet        string
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal().Msgf("error while loading env vars: %s ", err.Error())
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logWriter := zerolog.ConsoleWriter{
		Out:     colorable.NewColorableStdout(),
		NoColor: false,
	}
	log.Logger = log.Output(logWriter)

	bot := newBot()
	err = bot.Open()
	if err != nil {
		log.Fatal().Msgf("error while connecting to Discord: %s ", err.Error())
	}

	err = bot.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatal().Msg("unable to set trusted proxy")
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	bot.Use(cors.New(corsConfig))

	log.Info().Msg("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := bot.ApplicationCommandCreate(bot.State.User.ID, "", v)
		if err != nil {
			log.Panic().Msgf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	bot.AddHandler(bot.interactionHandler)

	bot.Engine.POST("/connect", bot.connectMiddleware, bot.connectionHandler)
	bind := os.Getenv("BIND")
	if bind == "" {
		log.Fatal().Msg("BIND must be set")
	}
	srv := &http.Server{
		Addr:    bind,
		Handler: bot.Engine,
	}

	go bot.checkAwaiting()

	idleConnsClosed := make(chan struct{})
	log.Info().Msg("Bot is now running. Press CTRL-C to exit.")
	go func() {
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
		<-sc
		log.Info().Msg("Shutting down server...")

		// The context is used to inform the server it has 5 seconds to finish
		// the request it is currently handling
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(ctx); err != nil {
			// Error from closing listeners, or context timeout:
			log.Fatal().Err(err).Msg("Server forced to shutdown")
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal().Err(err).Msg("Failed to start server")
	}

	<-idleConnsClosed
	_ = bot.Session.Close()

	log.Info().Msg("Removing commands...")
	for _, v := range registeredCommands {
		err := bot.ApplicationCommandDelete(bot.State.User.ID, "", v.ID)
		if err != nil {
			log.Fatal().Err(err).Msgf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}

	log.Info().Msg("Server exiting")
}

func newBot() Bot {
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatal().Msgf("error while instantiating Discord bot : %s ", err.Error())
	}

	bot := Bot{
		Session: dg,
		Engine:  gin.New(),
		awaitingConnectionMessageToUser: &awaitingConnectionMessageToUser{
			make(awaitingConnectionMessageToUserMap),
			&sync.RWMutex{},
		},

		userToAddress: &userToAddress{
			make(userToAddressMap),
			&sync.RWMutex{},
		},

		roleToContracts: &roleToContracts{
			make(roleToContractsMap),
			&sync.RWMutex{},
		},

		guildIDtoRoles: &guildIDtoRoles{
			make(guildIDtoRolesMap),
			&sync.RWMutex{},
		},
		contracts: &contracts{
			make(contractsMap),
			&sync.RWMutex{},
		},
		guilds: &guilds{
			make(guildsMap),
			&sync.RWMutex{},
		},

		database: newDB(),
	}
	err = bot.getAwaitingConnections(bot.awaitingConnectionMessageToUser.m)
	if err != nil {
		log.Fatal().Msgf("error while getting awaiting connections from db: %s ", err.Error())
	}
	err = bot.getUserAddress(bot.userToAddress.m)
	if err != nil {
		log.Fatal().Msgf("error while getting users addresses from db: %s ", err.Error())
	}
	err = bot.getGuildAndRoles(bot.guildIDtoRoles.m)
	if err != nil {
		log.Fatal().Msgf("error while getting guilds and roles from db: %s ", err.Error())
	}
	err = bot.getContracts(bot.contracts.m)
	if err != nil {
		log.Fatal().Msgf("error while getting contracts from db: %s ", err.Error())
	}
	err = bot.getGuilds(bot.guilds.m)
	if err != nil {
		log.Fatal().Msgf("error while getting guilds from db: %s ", err.Error())
	}
	err = bot.getRolesAndContracts(bot.roleToContracts.m)
	if err != nil {
		log.Fatal().Msgf("error while getting roles and contracts link from db: %s ", err.Error())
	}

	return bot

}

func (b *Bot) connectionHandler(c *gin.Context) {
	var ConnectData connectBind
	err := c.ShouldBind(&ConnectData)
	if err != nil {
		_ = c.Error(err)
		c.Status(http.StatusBadRequest)
		c.Abort()
		return
	}
	//fmt.Println(ConnectData)

	userInfo, ok := b.awaitingConnectionMessageToUser.m[ConnectData.Message]
	if !ok {
		_ = c.Error(errors.New("no awaiting for this message."))
		c.Status(http.StatusBadRequest)
		c.Abort()
		return
	}

	e, _ := base64.StdEncoding.DecodeString(ConnectData.Message)
	if !b.verifySig(ConnectData.Wallet, ConnectData.SignedMessage, e) {
		_ = c.Error(errors.New("can not verify signer."))
		c.Status(http.StatusBadRequest)
		c.Abort()
		return
	}

	b.addUserWallet(userInfo.userID, strings.ToLower(ConnectData.Wallet))

	b.guildIDtoRoles.RLock()
	_, ok = b.guildIDtoRoles.m[userInfo.guildID]
	b.guildIDtoRoles.RUnlock()
	if ok {
		go b.checkRoles(userInfo.userID, userInfo.guildID, strings.ToLower(ConnectData.Wallet), userInfo.interaction, ConnectData.Message)
	} else {
		msg := fmt.Sprintf("You succesfully connected this address:\n`%s`", strings.ToLower(ConnectData.Wallet))
		_, e := b.Session.FollowupMessageCreate(userInfo.interaction, false, &discordgo.WebhookParams{
			Content: msg,
			Flags:   64,
		})
		if e != nil {
			log.Err(e).Msg("error while sending followup message of connected address.")
		}
	}

	c.Status(200)
}

func (b *Bot) checkPerms(perms int64) bool {
	if perms&8 != 0 {
		return true
	}
	return false
}

func (b *Bot) connectMiddleware(c *gin.Context) {
	if len(b.awaitingConnectionMessageToUser.m) == 0 {
		_ = c.Error(errors.New("no connection awaiting"))
		c.Status(http.StatusTooEarly)
		c.Abort()
		return
	}
}

func (b *Bot) checkAwaiting() {
	t := time.NewTicker(time.Minute * 20)
	for {
		<-t.C
		b.awaitingConnectionMessageToUser.Lock()
		for k, u := range b.awaitingConnectionMessageToUser.m {
			if u.Time.Add(time.Minute * 20).Before(time.Now()) {
				delete(b.awaitingConnectionMessageToUser.m, k)
			}
		}
		b.awaitingConnectionMessageToUser.Unlock()
		_, err := b.NewDelete().
			Model((*awaitingConnectionDB)(nil)).
			Where("\"createdAt\" < ?", time.Now().Add(time.Minute*-20)).
			Exec(context.Background())

		if err != nil {
			log.Error().Msg(err.Error())
		}
	}
}
