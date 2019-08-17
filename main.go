package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// ConfigClient struct for client hash
type ConfigClient struct {
	DNS  string
	Name string
}

// Config top level struct for toml parser
type Config struct {
	Clients map[string]ConfigClient
}

func parseConfig(configPath string) Config {
	var config Config

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalln("Config file does not exists")
	} else {
		if configBlob, err := ioutil.ReadFile(configPath); err != nil {
			log.Fatalln(err)
		} else {
			if _, err := toml.Decode(string(configBlob), &config); err != nil {
				log.Fatalln("Config file is not valid", err)
			}
		}
	}
	return config
}

func main() {
	configPath := flag.String("config", "config.toml", "Path to TOML config")

	flag.Parse()

	config := parseConfig(*configPath)

	for token, data := range config.Clients {
		fmt.Println("Client name:", token, "Client token:", data.DNS)
	}
}
