package main

import (
	"encoding/base64"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"strconv"
	"strings"
	"time"
)

type messageComponents []discordgo.MessageComponent

var (
	integerOptionMinValue = float64(721)
	commands              = []*discordgo.ApplicationCommand{
		{
			Name:        "add-role-token",
			Description: "Command for linking a token to a role",
			Options: []*discordgo.ApplicationCommandOption{

				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "token-address",
					Description: "token address",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionInteger,
					Name:        "token-standard",
					Description: "token standard (721 or 1155)",
					MinValue:    &integerOptionMinValue,
					MaxValue:    1155,
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionRole,
					Name:        "role",
					Description: "Role option",
					Required:    true,
				},
			},
		},
		{
			Name:        "define-start-channel",
			Description: "Command for defining a start channel",
			Options: []*discordgo.ApplicationCommandOption{

				{
					Type:        discordgo.ApplicationCommandOptionChannel,
					Name:        "channel",
					Description: "Channel option",
					// Channel type mask
					ChannelTypes: []discordgo.ChannelType{
						discordgo.ChannelTypeGuildText,
					},
					Required: true,
				},
			},
		},
		{
			Name:        "show-roles-link",
			Description: "show roles and tokens link",
		},
	}
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate, b *Bot){
		"define-start-channel": func(s *discordgo.Session, i *discordgo.InteractionCreate, b *Bot) {
			if isAdmin := b.checkPerms(i.Member.Permissions); !isAdmin {
				return
			}
			options := i.ApplicationCommandData().Options
			err := b.checkGuild(i.GuildID)
			if err != nil {
				log.Err(err).Msg("error while inserting guild into db")
				return
			}
			/*
				_, err = s.ChannelMessageSendComplex(options[0].ChannelValue(s).ID, &discordgo.MessageSend{
					Content: "## Verify your assets \n" +
						"- We will never ask for any transaction. \n" +
						"- We will never ask for your seed phrase. \n" +
						"- We will never DM you.",
					Components: messageComponents{discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{discordgo.Button{
							Label:    "Start now!",
							Style:    discordgo.PrimaryButton,
							CustomID: "start",
						},
						},
					},
					},
				})

			*/
			_, err = s.ChannelMessageSendComplex(options[0].ChannelValue(s).ID, &discordgo.MessageSend{Embeds: []*discordgo.MessageEmbed{&discordgo.MessageEmbed{
				Description: "## Verify your assets \n" +
					"- We will never ask for any transaction. \n" +
					"- We will never ask for your seed phrase. \n" +
					"- We will never DM you.",
				Color: 4989874,
				Thumbnail: &discordgo.MessageEmbedThumbnail{
					URL: "https://cdn.discordapp.com/attachments/1097360869362499625/1111286348792463392/Screenshot_7.png",
				},
			},
			},
				Components: messageComponents{discordgo.ActionsRow{[]discordgo.MessageComponent{discordgo.Button{
					Label:    "Start now!",
					Style:    discordgo.PrimaryButton,
					CustomID: "start",
				},
				}}}})

			if err != nil {
				log.Err(err).Msg("error while sending start message")
				err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "start message failed to create.",
						Flags:   64,
					},
				})
				return
			}

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "start message sent successfully.",
					Flags:   64,
				},
			})
			if err != nil {
				log.Err(err).Msg("error while reacting to define-start-channel command")
			}

		},
		"add-role-token": func(s *discordgo.Session, i *discordgo.InteractionCreate, b *Bot) {
			if isAdmin := b.checkPerms(i.Member.Permissions); !isAdmin {
				return
			}
			options := i.ApplicationCommandData().Options
			err := b.checkGuild(i.GuildID)
			if err != nil {
				log.Err(err).Msg("error while inserting guild into db")
				return
			}
			for _, c := range b.roleToContracts.m[options[2].RoleValue(s, i.GuildID).ID] {
				co := contract{
					Address:   strings.ToLower(options[0].StringValue()),
					tokenType: tokenType(options[1].IntValue()),
				}
				if c == co {
					err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "Link was connected already.",
							Flags:   64,
						},
					})
					if err != nil {
						log.Err(err).Msg("error while sending link confirmation message.")
					}
					return
				}
			}

			err = b.addRoleContract(contract{
				Address:   strings.ToLower(options[0].StringValue()),
				tokenType: tokenType(options[1].IntValue()),
			},
				i.GuildID,
				options[2].RoleValue(s, i.GuildID).ID)

			if err != nil {
				err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Error: Link was not created.",
						Flags:   64,
					},
				})
				if err != nil {
					log.Err(err).Msg("error while sending link failed message.")
				}
				return
			}

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Link created successfully.",
					Flags:   64,
				},
			})
			if err != nil {
				log.Err(err).Msg("error while sending link confirmation message.")
			}

		},
		"show-roles-link": func(s *discordgo.Session, i *discordgo.InteractionCreate, b *Bot) {
			if isAdmin := b.checkPerms(i.Member.Permissions); !isAdmin {
				return
			}
			b.guildIDtoRoles.RLock()
			b.roleToContracts.RLock()
			for _, role := range b.guildIDtoRoles.m[i.GuildID] {
				var msg string
				for _, contract := range b.roleToContracts.m[role] {
					msg += fmt.Sprintf("role: <@&%s>, token: %s, standard: %v \n", role, contract.Address, contract.tokenType)
				}
				err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: msg,
						Flags:   64,
					},
				})
				if err != nil {
					log.Err(err).Msg("error while sending roles link response")
				}
			}
			b.roleToContracts.RUnlock()
			b.guildIDtoRoles.RUnlock()
		},
	}
)

func (b *Bot) addRoleContract(c contract, gID, rID string) error {
	var err error
	intGID, _ := strconv.Atoi(gID)
	intRID, _ := strconv.Atoi(rID)
	t := time.Now()

	b.roleToContracts.RLock()
	_, ok := b.roleToContracts.m[rID]
	b.roleToContracts.RUnlock()
	if !ok {
		err := b.insertRole(role{
			ID:        &intRID,
			GuildID:   &intGID,
			CreatedAt: &t,
		})
		if err != nil {
			log.Err(err).Msg("error while inserting role into db")
			return err
		}

	}
	var contractID *int
	var contract contractDB
	tkType := int(c.tokenType)
	nID := 1
	b.contracts.RLock()
	_, ok = b.contracts.m[fmt.Sprintf("%s_1", c.Address)]
	b.contracts.RUnlock()
	if !ok {
		contract = contractDB{
			Address:   c.Address,
			Type:      &tkType,
			NetworkID: &nID,
			CreatedAt: &t,
		}
		contractID, err = b.insertContract(contract)
		if err != nil {
			log.Err(err).Msg("error while inserting contract into db")
			return err
		}
		contract.ID = contractID
	}

	err = b.insertRoleContract(roleContract{
		RoleID:     &intRID,
		ContractID: contractID,
	})
	if err != nil {
		log.Err(err).Msg("error while inserting role-contract into db")
		return err
	}
	b.contracts.Lock()
	b.contracts.m[fmt.Sprintf("%s_1", c.Address)] = contract
	b.contracts.Unlock()
	b.roleToContracts.Lock() // TODO MOVE THIS TO UP AFTER ADDING DELETE FUNCTION
	b.roleToContracts.m[rID] = append(b.roleToContracts.m[rID], c)
	b.roleToContracts.Unlock()
	b.guildIDtoRoles.Lock()
	b.guildIDtoRoles.m[gID] = append(b.guildIDtoRoles.m[gID], rID)
	b.guildIDtoRoles.Unlock()
	return nil
}

func (b *Bot) interactionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	mC, ok := i.Interaction.Data.(discordgo.MessageComponentInteractionData)
	if !ok {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i, b)
		}
		return
	}
	switch strings.Split(mC.CustomID, "_")[0] {
	case "start":
		b.userToAddress.RLock()
		adds, ok := b.userToAddress.m[i.Member.User.ID]
		b.userToAddress.RUnlock()
		if !ok {
			t := time.Now()
			intUID, _ := strconv.Atoi(i.Member.User.ID)
			u, _ := b.User(i.Member.User.ID)
			err := b.insertUser(
				user{
					ID:        &intUID,
					Name:      u.Username,
					CreatedAt: &t,
				})
			if err != nil {
				log.Err(err).Msg("error while inserting user into db")
				return
			}
		}

		var components []discordgo.MessageComponent
		var message string
		if len(adds) != 0 {
			buttons, me := b.generateButtonsForUser(adds)
			message = me
			for _, bu := range buttons {
				components = append(components, bu)
			}
		} else {
			message = "`no address connected yet.`\n"
		}
		components = append(components, b.newButton("new address", "new-address", "", discordgo.PrimaryButton))
		//component = append(component, b.newButton("delete address", "delete-address", "", discordgo.PrimaryButton)) //TODO
		err := s.InteractionRespond(i.Interaction, b.newInteraction("Your Connected Addresses", fmt.Sprintf(
			"# Connected Addresses:\n%s\n", message), components, "we only support metamask for now.", 4))
		if err != nil {
			log.Err(err).Msg("error while sending start message")
		}

	case "new-address":
		b.newWalletInteraction(s, i)

	case "use-address":
		b.guildIDtoRoles.RLock()
		_, ok = b.guildIDtoRoles.m[i.GuildID]
		b.guildIDtoRoles.RUnlock()
		if !ok {
			e := b.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseType(4),
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("Nothing to unlock for now."),
					Title:   "unlock",
					Flags:   64,
				},
			})
			if e != nil {
				log.Err(e).Msg("error while sending unlocked roles response.")
			}
			return
		}
		addressIndex, _ := strconv.Atoi(strings.Split(mC.CustomID, "_")[1])
		b.userToAddress.RLock()
		address := b.userToAddress.m[i.Member.User.ID][addressIndex]
		b.userToAddress.RUnlock()
		b.checkRoles(i.Member.User.ID, i.GuildID, address, i.Interaction, "")

	case "delete-address":

		//TODO mark deleted on db
	}

}

func (b *Bot) generateButtonsForUser(addresses []string) ([]discordgo.Button, string) {
	var buttons []discordgo.Button
	var message string
	for i, a := range addresses {
		buttons = append(buttons, b.newButton(fmt.Sprintf("%v", i+1), fmt.Sprintf("use-address_%v", i), "", discordgo.PrimaryButton))
		message += fmt.Sprintf("%v: `%s`\n", i+1, a)
	}
	return buttons, message
}

func (b *Bot) newInteraction(title, content string, mC []discordgo.MessageComponent, footer string, respType int) *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseType(respType),
		Data: &discordgo.InteractionResponseData{
			Title: title,
			Flags: 64,
			Embeds: []*discordgo.MessageEmbed{&discordgo.MessageEmbed{
				Description: content,
				Color:       4989874,
				Footer:      &discordgo.MessageEmbedFooter{Text: footer},
			}},
			Components: messageComponents{discordgo.ActionsRow{
				Components: mC,
			},
			},
		},
	}
}

func (b *Bot) newButton(label, customID, url string, style discordgo.ButtonStyle) discordgo.Button {
	return discordgo.Button{
		Label:    label,
		Style:    style,
		CustomID: customID,
		URL:      url,
	}
}

func (b *Bot) newWalletInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	g, _ := b.Guild(i.GuildID)
	message := fmt.Sprintf("user id: %s \nguild id: %s\n%s", i.Member.User.ID, i.GuildID, g.Name)
	encodedMessage := base64.StdEncoding.EncodeToString([]byte(message))
	err := s.InteractionRespond(i.Interaction, b.newInteraction("", fmt.Sprintf("### You should expect to sign the following message when prompted by MetaMask:  \n ```%s``` \n no transaction will be necessary.", message),
		[]discordgo.MessageComponent{
			b.newButton("Connect Wallet", "", fmt.Sprintf("http://localhost:3000/connect?sign=%s", encodedMessage), discordgo.LinkButton),
		}, "", 4))
	if err != nil {
		log.Err(err).Msg("error while sending new wallet interaction")
	}

	b.awaitingConnectionMessageToUser.RLock()
	_, ok := b.awaitingConnectionMessageToUser.m[encodedMessage]
	b.awaitingConnectionMessageToUser.RUnlock()
	if !ok {
		t := time.Now()
		intUID, _ := strconv.Atoi(i.Member.User.ID)
		b.awaitingConnectionMessageToUser.Lock()
		b.awaitingConnectionMessageToUser.m[encodedMessage] = &userInfo{
			userID:      i.Member.User.ID,
			guildID:     i.GuildID,
			interaction: i.Interaction,
			Time:        &t,
		}
		b.awaitingConnectionMessageToUser.Unlock()

		intGID, _ := strconv.Atoi(i.GuildID)
		IntID, _ := strconv.Atoi(i.Interaction.ID)
		intTOK, _ := strconv.Atoi(i.Interaction.Token)
		intAPP, _ := strconv.Atoi(i.Interaction.AppID)
		err := b.insertAwaitingConnection(awaitingConnectionDB{
			Message:          encodedMessage,
			UserID:           &intUID,
			GuildID:          &intGID,
			InteractionID:    &IntID,
			InteractionToken: &intTOK,
			InteractionAppID: &intAPP,
			CreatedAt:        &t,
		})
		if err != nil {
			log.Err(err).Msg("error while inserting awaiting connection into db")
		}
	}

	return
}

func (b *Bot) checkGuild(id string) error {
	b.guilds.Lock()
	defer b.guilds.Unlock()
	_, ok := b.guilds.m[id]
	if !ok {
		t := time.Now()
		intGID, _ := strconv.Atoi(id)
		g, _ := b.Session.Guild(id)
		guild := guild{
			ID:        &intGID,
			Name:      g.Name,
			CreatedAt: &t,
		}
		err := b.insertGuild(guild)
		if err != nil {
			return err
		}
		b.guilds.m[id] = guild
		return nil
	}
	return nil
}

//TODO check returns in errors

// if role reaches no contracts role will be detached from guild and deleted on db
// only roles get deleted from db
