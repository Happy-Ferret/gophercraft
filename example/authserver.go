package main

import (
	"log"

	"github.com/superp00t/gophercraft/authserver"
	"github.com/superp00t/gophercraft/packet"
	"github.com/superp00t/gophercraft/srp"
)

func GetAccount(user string) *authserver.Account {
	u := "USER"
	p := "PASSWORD"
	if user != u {
		return nil
	}

	return &authserver.Account{
		Username:     u,
		IdentityHash: srp.HashCredentials(u, p),
	}
}

func GetRealms(u string) []packet.RealmListing {
	return []packet.RealmListing{
		{
			Type:       1,
			Color:      packet.REALM_GREEN,
			Locked:     false,
			Flags:      0,
			Name:       "Gophercraft",
			Address:    "0.0.0.0:8085",
			Population: 2,
			Characters: 0,
			Timezone:   8,
			ID:         44,
		},
		{
			Type:       1,
			Color:      packet.REALM_YELLOW,
			Locked:     false,
			Flags:      0,
			Name:       "The WoW server",
			Address:    "0.0.0.0:8085",
			Population: 2,
			Characters: 0,
			Timezone:   8,
			ID:         44,
		},
		{
			Type:       1,
			Color:      packet.REALM_RED,
			Locked:     false,
			Flags:      0,
			Name:       "Written in Go!",
			Address:    "0.0.0.0:8085",
			Population: 2,
			Characters: 0,
			Timezone:   8,
			ID:         44,
		},
	}
}

func main() {
	cfg := &authserver.Config{
		Listen:     "localhost:3724",
		GetAccount: GetAccount,
		ListRealms: GetRealms,
	}

	log.Fatal(authserver.Start(cfg))
}
