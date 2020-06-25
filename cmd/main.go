package main

import (
	"flag"

	"github.com/gredinger/farmer"
)

func main() {
	var fn string
	flag.StringVar(&fn, "c", "server.ini", "Configuration file")
	flag.Parse()
	webApp := farmer.WebApp{}
	webApp.LoadSettings(fn)
	webApp.Run()
}
