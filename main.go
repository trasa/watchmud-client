package main

import (
	"flag"
	"github.com/trasa/watchmud-message"
	"log"
	"os"
	"os/signal"
)

func main() {
	// TODO read from yaml configuration or something

	playerName := flag.String("player", "somedood", "player name")
	host := flag.String("host", "localhost", "server host name")
	port := flag.Int("port", 10000, "server port")
	flag.Parse()

	// connect client
	client, err := Connect(*host, *port)
	if err != nil {
		log.Fatal("Failed to connect", err)
	}
	signal.Notify(client.quitSignal, os.Interrupt)

	// send login request
	password := "NotImplemented"

	loginReq := message.LoginRequest{
		PlayerName: *playerName,
		Password:   password,
	}
	loginMsg, err := message.NewGameMessage(loginReq)
	if err != nil {
		log.Fatalf("Error creating login message: %v", err)
	}
	client.SendMessage(loginMsg)

	client.readStdin()
}
