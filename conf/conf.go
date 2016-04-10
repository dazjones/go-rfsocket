package conf

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Broker    string
	Hostname  string
	MySQLUser string
	MySQLPass string
	MySQLHost string
	MySQLName string
	MySQLPort string
}

func Init(c *Configuration) {
	file, err := os.Open("config.conf")

	if err != nil {
		fmt.Println("Configuration File Error: ", err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)

	if err != nil {
		fmt.Println("Configuration Decode Error: ", err)
	}
}
