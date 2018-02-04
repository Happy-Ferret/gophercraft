package main

import (
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/ogier/pflag"
	"github.com/superp00t/gophercraft/client"
)

var (
	username   = pflag.StringP("username", "u", "username", "username")
	pass       = pflag.StringP("pass", "p", "password", "password")
	authServer = pflag.StringP("authserver", "a", "localhost:3724", "authserver")
	realm      = pflag.StringP("realm", "r", "Icecrown", "realm")
	player     = pflag.StringP("player", "y", "wew", "playername")
)

func main() {
	pflag.Parse()
	/*	c, err := client.AuthConnect(&client.Config{
			Username:   "fuck",
			Password:   "puck",
			AuthServer: "localhost:3724",
		})
	*/
	c, err := client.AuthConnect(&client.Config{
		Username:   *username,
		Password:   *pass,
		AuthServer: *authServer,
		Playername: *player,
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println(spew.Sdump(c.RealmList))

	for _, ve := range c.RealmList.Realms {
		if ve.Name == *realm {
			err := c.WorldConnect(ve.Address)
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(1000000 * time.Minute)
		}
	}
}
