package main

import (
	"flag"
	"github.com/dazjones/go-rfsocket/conf"
	"github.com/dazjones/go-rfsocket/models"
	"log"
)

var config conf.Configuration

func main() {
	sender_bin := flag.String("sender", "/usr/bin/rfsender", "Path to RF sender binary")

	flag.Parse()

	if *sender_bin == "" {
		log.Fatalf("Flag sender is empty")
	}

	conf.Init(&config)
	models.Init(&config)

}
