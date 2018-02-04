package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ogier/pflag"
	"github.com/superp00t/gophercraft/authserver"
	"github.com/superp00t/gophercraft/gcore"
	log "github.com/superp00t/gophercraft/gcore/glogger"
)

var (
	aaddr  = pflag.StringP("auth_listen", "a", ":3724", "the IP address to serve the Authserver on")
	haddr  = pflag.StringP("http_listen", "l", ":8086", "the IP address to serve the JSON API on")
	driver = pflag.StringP("driver", "d", "mysql", "the XORM driver to use for the database")
	source = pflag.StringP("source", "s", "root:root@/gcraft_core", "the XORM driver to use for the database")
)

func main() {
	pflag.Parse()

	gcore.PrintLicense()

	log.Println("Starting Gophercraft Core Auth Server...")
	log.Println("~ ~ ~")

	core, err := gcore.NewCore(*driver, *source)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database opened without issue.")

	go func() {
		log.Println("Starting HTTP server at", *haddr)
		mux := core.WebAPI()
		log.Fatal(http.ListenAndServe(*haddr, mux))
	}()

	log.Println("Starting Authserver at", *aaddr)
	log.Fatal(authserver.Start(&authserver.Config{
		Listen:     *aaddr,
		GetAccount: core.GetAccount,
		ListRealms: core.ListRealms,
		StoreKey:   core.StoreKey,
	}))
}
