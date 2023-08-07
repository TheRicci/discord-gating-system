package main

import (
	"github.com/uptrace/bun"
	"time"
)

type user struct {
	ID            *int
	Name          string
	CreatedAt     *time.Time `bun:"createdAt"`
	bun.BaseModel `bun:"gatingDB_user"`
}

type wallet struct {
	ID            *int
	Address       string
	NetworkID     *int
	UserID        *int
	CreatedAt     *time.Time `bun:"createdAt"`
	bun.BaseModel `bun:"gatingDB_wallet"`
}

type guild struct {
	ID            *int
	Name          string
	CreatedAt     *time.Time `bun:"createdAt"`
	bun.BaseModel `bun:"gatingDB_guild"`
}

type role struct {
	ID            *int
	GuildID       *int
	CreatedAt     *time.Time `bun:"createdAt"`
	Contracts     []string   `bun:"-"`
	bun.BaseModel `bun:"gatingDB_role"`
}

type contractDB struct {
	ID            *int
	Address       string
	Type          *int
	NetworkID     *int
	CreatedAt     *time.Time `bun:"createdAt"`
	bun.BaseModel `bun:"gatingDB_contract"`
}

type network struct {
	ChainID       *int
	Name          string
	HttpRPC       string     `bun:"httpRPC"`
	CreatedAt     *time.Time `bun:"createdAt"`
	bun.BaseModel `bun:"gatingDB_network"`
}

type getRoleContract struct {
	ID              *int
	RoleID          *int
	ContractID      *int
	ContractAddress string `bun:"contract_address"`
	ContractType    *int   `bun:"contract_type"`
	bun.BaseModel   `bun:"gatingDB_contract_role"`
}

type roleContract struct {
	ID            *int
	RoleID        *int
	ContractID    *int
	bun.BaseModel `bun:"gatingDB_contract_role"`
}

type awaitingConnectionDB struct {
	Message          string
	UserID           *int
	GuildID          *int
	InteractionID    *int       `bun:"interactionID"`
	InteractionToken *int       `bun:"interactionToken"`
	InteractionAppID *int       `bun:"interactionAppID"`
	CreatedAt        *time.Time `bun:"createdAt"`
	bun.BaseModel    `bun:"gatingDB_awaitingconnection"`
}
