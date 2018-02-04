package main

import (
	"github.com/superp00t/gophercraft/gcore"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ogier/pflag"
	log "github.com/superp00t/gophercraft/gcore/glogger"
	"github.com/superp00t/gophercraft/worldserver"
)

var (
	waddr     = pflag.StringP("world_listen", "l", ":8085", "the IP address to serve the World Server on")
	realmName = pflag.StringP("realm_name", "r", "demo", "name of the realm server. MUST be unique. Synced with database.")
	ipAddr    = pflag.StringP("ip_address", "i", "ip", "public IP of the realm server. Synced with database.")
	endpoint  = pflag.StringP("endpoint", "e", "http://localhost:8086", "Gophercraft Core HTTP API server")
	apiKey    = pflag.StringP("api_key", "a", "_proper_", "authentication key for using Gophercraft Core API")
	driver    = pflag.StringP("driver", "d", "mysql", "the XORM driver to use for the database")
	source    = pflag.StringP("source", "s", "root:root@/gcraft_core", "the XORM driver to use for the database")
)

func main() {
	pflag.Parse()

	gcore.PrintLicense()

	if *realmName == "" {
		log.Fatal("You must choose a unique realm name to add to the list: `gcraft_core_world -r [realm name]`")
	}

	log.Fatal(worldserver.Start(&worldserver.Config{
		Listen:        *waddr,
		RealmName:     *realmName,
		Driver:        *driver,
		Source:        *source,
		PublicAddress: *ipAddr,
		WardenEnabled: false,
	}))
}
