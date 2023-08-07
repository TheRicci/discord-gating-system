package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"os"
	"strings"
)

type database struct {
	*bun.DB
}

func newDB() *database {
	var connector *pgdriver.Connector

	if !strings.EqualFold(os.Getenv("DATABASE_URL"), "") {
		dsn := os.Getenv("DATABASE_URL")
		connector = pgdriver.NewConnector(pgdriver.WithDSN(dsn))
	} else {
		connector = pgdriver.NewConnector(
			pgdriver.WithAddr(os.Getenv("DB_ADDR")),
			pgdriver.WithUser(os.Getenv("DB_USER")),
			pgdriver.WithPassword(os.Getenv("DB_PASS")),
			pgdriver.WithDatabase(os.Getenv("DB_NAME")),
		)
	}

	sqldb := sql.OpenDB(connector)
	db := bun.NewDB(sqldb, pgdialect.New(), bun.WithDiscardUnknownColumns())
	if db == nil {
		log.Fatal().Err(errors.New("failed to connect db"))
	}

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true), // log everything
	))

	log.Info().Msg("db connected")
	return &database{
		db,
	}
}

func (db *database) getAwaitingConnections(m awaitingConnectionMessageToUserMap) error {
	var acs []awaitingConnectionDB

	err := db.NewSelect().Model(&acs).Scan(context.Background())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return err
	}
	for _, ac := range acs {
		m[ac.Message] = &userInfo{
			userID:  fmt.Sprintf("%v", *ac.UserID),
			guildID: fmt.Sprintf("%v", *ac.GuildID),
			interaction: &discordgo.Interaction{
				ID:    fmt.Sprintf("%v", *ac.InteractionID),
				Token: fmt.Sprintf("%v", *ac.InteractionToken),
				AppID: fmt.Sprintf("%v", *ac.InteractionAppID),
			},
			Time: ac.CreatedAt,
		}
	}
	return nil
}

func (db *database) getUserAddress(addressMap userToAddressMap) error {
	var wallets []wallet

	err := db.NewSelect().Model(&wallets).Scan(context.Background())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return err
	}

	for _, w := range wallets {
		addressMap[fmt.Sprintf("%v", *w.UserID)] = append(addressMap[fmt.Sprintf("%v", *w.UserID)], w.Address)
	}
	return nil
}

func (db *database) getGuildAndRoles(m guildIDtoRolesMap) error {
	var roles []role

	err := db.NewSelect().Model(&roles).Scan(context.Background())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return err
	}
	for _, r := range roles {
		m[fmt.Sprintf("%v", *r.GuildID)] = append(m[fmt.Sprintf("%v", *r.GuildID)], fmt.Sprintf("%v", *r.ID))
	}
	return nil
}

func (db *database) getRolesAndContracts(m roleToContractsMap) error {
	var rolesAndContracts []getRoleContract

	err := db.NewSelect().Model(&rolesAndContracts).
		ColumnExpr("role_id").
		ColumnExpr("c.address AS contract_address, c.type AS contract_type").
		Join("JOIN \"gatingDB_contract\" AS c ON c.id = contract_id").
		Scan(context.Background())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil
		}
		return err
	}

	for _, rC := range rolesAndContracts {
		m[fmt.Sprintf("%v", *rC.RoleID)] = append(m[fmt.Sprintf("%v", *rC.RoleID)], contract{
			Address:   rC.ContractAddress,
			tokenType: tokenType(*rC.ContractType),
		})
	}
	return nil
}

func (db *database) getContracts(m contractsMap) error {
	var cS []contractDB
	err := db.NewSelect().Model(&cS).Scan(context.Background())
	if err != nil {
		return err
	}

	for _, c := range cS {
		m[fmt.Sprintf("%s_%v", c.Address, *c.NetworkID)] = c
	}

	return nil
}

func (db *database) getGuilds(m guildsMap) error {
	var gS []guild
	err := db.NewSelect().Model(&gS).Scan(context.Background())
	if err != nil {
		return err
	}

	for _, g := range gS {
		m[fmt.Sprintf("%v", *g.ID)] = g
	}

	return nil
}

func (db *database) insertUser(u user) error {
	_, err := db.NewInsert().Model(&u).Exec(context.Background())
	if err != nil {
		if db.checkIfUniqueConstraintPK(err) {
			return nil
		}
		return err
	}
	return nil
}

func (db *database) insertWallet(w wallet) error {
	_, err := db.NewInsert().Model(&w).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (db *database) insertGuild(g guild) error {
	_, err := db.NewInsert().Model(&g).Exec(context.Background())
	if err != nil {
		if db.checkIfUniqueConstraintPK(err) {
			return nil
		}
		return err
	}
	return nil
}

func (db *database) insertRole(r role) error {
	_, err := db.NewInsert().Model(&r).Exec(context.Background())
	if err != nil {
		if db.checkIfUniqueConstraintPK(err) {
			return nil
		}
		return err
	}
	return nil
}

func (db *database) insertContract(c contractDB) (*int, error) {
	var id int
	_, err := db.NewInsert().Model(&c).Returning("ID").Exec(context.Background(), &id)
	if err != nil {
		if db.checkIfUniqueConstraintPK(err) {
			err = db.NewSelect().Model(&c).Where("address like ?", c.Address).Scan(context.Background())
			if err != nil {
				return nil, err
			}
			return c.ID, nil
		}
		return nil, err
	}
	return &id, nil
}

func (db *database) insertRoleContract(rC roleContract) error {
	_, err := db.NewInsert().Model(&rC).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (db *database) insertAwaitingConnection(connectionDB awaitingConnectionDB) error {
	_, err := db.NewInsert().Model(&connectionDB).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (db *database) deleteAwaitingConnection(m string) error {
	_, err := db.NewDelete().Model((*awaitingConnectionDB)(nil)).Where("message like ?", m).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (db *database) checkIfUniqueConstraintPK(err error) bool {
	if e, ok := err.(pgdriver.Error); ok {
		if e.Field('C') == "23505" {
			return true
		}
		return false
	}
	return false
}
