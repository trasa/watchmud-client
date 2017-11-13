package main

import (
	"flag"
	"github.com/trasa/watchmud/message"
	"log"
	"os"
	"os/signal"
)

const SERVER_HOST = "localhost"
const SERVER_PORT = 10000

func main() {
	// TODO read from yaml configuration or something
	// TODO override with command line args

	playerName := flag.String("player", "somedood", "player name")
	flag.Parse()

	// connect client
	client, err := Connect(SERVER_HOST, SERVER_PORT)
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
