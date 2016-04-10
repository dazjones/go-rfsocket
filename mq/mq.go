package mq

import (
	"github.com/dazjones/go-rfsocket/conf"
	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	"log"
)

var f MQTT.MessageHandler = func(client *MQTT.Client, msg MQTT.Message) {
	log.Printf("Topic: %s, Payload: %s\n", msg.Topic(), msg.Payload())
}

func Init(config *conf.Configuration) {
	opts := MQTT.NewClientOptions().AddBroker(config.Broker)
	opts.SetClientID("go-rfsocket")
	opts.SetDefaultPublishHandler(f)

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := c.Subscribe("/#", 2, nil); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
}
