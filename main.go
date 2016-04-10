package main

import (
	"flag"
	"github.com/dazjones/go-rfsocket/conf"
	"github.com/dazjones/go-rfsocket/models"
	"github.com/dazjones/go-rfsocket/mq"
	"log"
)

var config conf.Configuration

func main() {
	sender_bin := flag.String("sender", "/usr/bin/rfsender", "Path to RF sender binary")

	flag.Parse()

	if *sender_bin == "" {
		log.Fatalf("Flag sender is empty")
	}

	log.Println("Initialising config");
	conf.Init(&config)

	log.Println("Initialising models");
	models.Init(&config)

	log.Println("Initialising MQTT");
	mq.Init(&config)

	log.Println("Ready");

	for {

	}

}
