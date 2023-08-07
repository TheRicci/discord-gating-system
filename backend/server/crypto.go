package main

import (
	"discord-gating-system/erc1155"
	"discord-gating-system/erc721"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
	"math/big"
	"strconv"
	"strings"
	"time"
)

type tokenType uint

const (
	ERC721  tokenType = 721
	ERC1155           = 1155
)

func (b *Bot) checkRoles(userID, guildID, address string, i *discordgo.Interaction, message string) {
	if message != "" {
		fmt.Println(i)
		b.awaitingConnectionMessageToUser.Lock()
		defer func() {
			err := b.deleteAwaitingConnection(message)
			if err != nil {
				log.Err(err).Msg("error while deleting awaiting connection in db")
			}
			delete(b.awaitingConnectionMessageToUser.m, message)
		}()
		defer b.awaitingConnectionMessageToUser.Unlock()
	}

	client, err := ethclient.Dial("https://eth.llamarpc.com")
	if err != nil {
		log.Err(err).Msg("error while connecting to ethereum node")
	}

	var passedRoles []string
	b.guildIDtoRoles.RLock()
	roles := b.guildIDtoRoles.m[guildID]
	b.guildIDtoRoles.RUnlock()

	for _, role := range roles {
		b.roleToContracts.RLock()
		cS := b.roleToContracts.m[role]
		b.roleToContracts.RUnlock()

		pass := true
		for _, contract := range cS {
			contractAddress := common.HexToAddress(contract.Address)
			var number *big.Int
			switch contract.tokenType {
			case ERC721:
				instance, err := erc721.NewToken721(contractAddress, client)
				if err != nil {
					log.Err(err).Msg("error while instantiating ERC721 token.")
					return
				}
				number, err = instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))
				if err != nil {
					log.Err(err).Msg("error while getting ERC721 token balance.")
					return
				}
			case ERC1155:
				instance, err := erc1155.NewToken1155(contractAddress, client)
				if err != nil {
					log.Err(err).Msg("error while instantiating ERC1155 token.")
					return
				}
				number, err = instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address), new(big.Int).SetInt64(0))
				if err != nil {
					log.Err(err).Msg("error while getting ERC1155 token balance.")
					return
				}
			}

			//fmt.Println(number)
			if number.Int64() == 0 {
				pass = false
				break
			}

		}

		if !pass {
			if message == "" {
				err := b.InteractionRespond(i, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseType(4),
					Data: &discordgo.InteractionResponseData{
						Content: "No roles to unlock with this wallet.",
						Flags:   64,
					},
				})
				if err != nil {
					log.Err(err).Msg("error while sending no unlocked roles response.")
				}
			} else {
				_, err := b.Session.FollowupMessageCreate(i, false, &discordgo.WebhookParams{
					Content: "No roles to unlock with this wallet.",
					Flags:   64,
				})
				if err != nil {
					log.Err(err).Msg("error while sending followup message of no unlocked roles.")
				}
			}
			return
		}

		passedRoles = append(passedRoles, role)
		err = b.Session.GuildMemberRoleAdd(guildID, userID, role)
		if err != nil {
			log.Err(err).Msgf("error while assigning role %s to user %s on guild %s", role, userID, role)
		}
	}

	msg := "Unlocked roles: \n"

	for _, role := range passedRoles {
		msg += fmt.Sprintf("<@&%s> \n", role)
	}

	if message == "" {
		e := b.InteractionRespond(i, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseType(4),
			Data: &discordgo.InteractionResponseData{
				Content: msg,
				Title:   "unlocked roles",
				Flags:   64,
			},
		})
		if e != nil {
			log.Err(e).Msg("error while sending unlocked roles response.")
		}
	} else {
		_, err = b.Session.FollowupMessageCreate(i, false, &discordgo.WebhookParams{
			Content: msg,
			Flags:   64,
		})
		if err != nil {
			log.Err(err).Msg("error while sending followup message of unlocked roles.")
		}
	}

}

func (b *Bot) addUserWallet(uID, address string) {
	intUID, _ := strconv.Atoi(uID)
	nID := 1
	t := time.Now()
	b.userToAddress.Lock()
	defer b.userToAddress.Unlock()
	adds, _ := b.userToAddress.m[uID]

	for _, a := range adds {
		if a == address {
			return
		}
	}
	err := b.insertWallet(wallet{
		UserID:    &intUID,
		Address:   address,
		NetworkID: &nID,
		CreatedAt: &t,
	})
	if err != nil {
		log.Err(err).Msg("error while inserting wallet into db")
		return
	}

	b.userToAddress.m[uID] = append(b.userToAddress.m[uID], address)
}

func (b *Bot) verifySig(from, sigHex string, msg []byte) bool {
	sig := hexutil.MustDecode(sigHex)

	msg = accounts.TextHash(msg)
	if sig[crypto.RecoveryIDOffset] == 27 || sig[crypto.RecoveryIDOffset] == 28 {
		sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}

	recovered, err := crypto.SigToPub(msg, sig)
	if err != nil {
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*recovered)

	return from == strings.ToLower(recoveredAddr.Hex())
}
