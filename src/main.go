package main

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"gotcpd"
)

func main() {
	config := loadConfiguration()
	gotcpd.AppConfig = config
	server := gotcpd.NewServer(config.Server.Host, config.Server.Port)

	log.Print("Starting listener")

	err := server.Listen()

	if err != nil {
		log.Fatalf("Cannot start listener: %s", err)
	}

	log.Printf("Listening on %s:%d", server.Host, server.Port)

	log.Print("Starting handler")

	connectionHandler := gotcpd.NewConnectionHandler(server.GetConnectionChannel())

	connectionHandler.Run()

	log.Print("Handler started")

	for {

	}
}

func loadConfiguration() (config gotcpd.Config) {
	data, err := ioutil.ReadFile("../config.json")

	if err != nil {
		log.Fatalf("Cannot load configuration file: %s", err)
	}

	config = gotcpd.Config{}

	err = json.Unmarshal(data, &config);

	if err != nil {
		log.Fatalf("Cannot parse json: %s", err)
	}

	return
}
