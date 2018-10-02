package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
)

func main() {
	// TODO read from yaml configuration or something

	host := flag.String("host", "localhost", "server host name")
	port := flag.Int("port", 10000, "server port")
	flag.Parse()

	ActiveConfig = &Config{
		serverHost: *host,
		serverPort: *port,
	}

	// connect client
	client, err := Connect(ActiveConfig.serverHost, ActiveConfig.serverPort)
	if err != nil {
		log.Fatal("Failed to connect", err)
	}
	signal.Notify(client.quitSignal, os.Interrupt)

	clientui := NewClientUI(client)
	// this runs ui event loop
	clientui.initUi()
}
