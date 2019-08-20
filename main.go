package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

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

var config Config

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

func handler(w http.ResponseWriter, r *http.Request) {
	if token, hasToken := r.Header["Access-Token"]; hasToken {
		if client, isClient := config.Clients[token[0]]; isClient {
			address := strings.Split(r.RemoteAddr, ":")[0]
			fmt.Fprintf(w, "Will generate new record %s:%s\n", client.DNS, address)
		} else {
			fmt.Fprintf(w, "No client found %s", token)
		}
	} else {
		fmt.Fprintf(w, "Go Away!")
	}
}

func main() {
	configPath := flag.String("config", "config.toml", "Path to TOML config")

	flag.Parse()

	config = parseConfig(*configPath)

	fmt.Println("Total clients ", len(config.Clients))
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
