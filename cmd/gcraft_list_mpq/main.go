package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ogier/pflag"
	"github.com/superp00t/gophercraft/blizzardry/mpq"
)

func main() {
	var fp = pflag.StringP("fp", "f", "...", "filepath to .MPQ file")
	pflag.Parse()

	m, err := mpq.Open(*fp)
	if err != nil {
		log.Fatal(err)
	}

	d, _ := json.MarshalIndent(m.ListFiles(), "", "  ")
	fmt.Println(string(d))
}
