package main

import (
	"flag"
	"github.com/trasa/watchmud/message"
	"log"
	"os"
	"os/signal"
)

const SERVER_ADDR = "ws://localhost:8888/ws"

func main() {
	// TODO read from yaml configuration or something
	// TODO override with command line args

	playerName := flag.String("player", "somedood", "player name")
	flag.Parse()

	// connect client
	client, err := Connect(SERVER_ADDR)
	if err != nil {
		log.Fatal("Failed to connect", err)
	}
	signal.Notify(client.quitSignal, os.Interrupt)

	// send login request
	password := "NotImplemented"

	loginReq := message.LoginRequest{
		Request:    message.RequestBase{MessageType: "login"},
		PlayerName: *playerName,
		Password:   password,
	}
	client.SendRequest(loginReq)

	client.readStdin()

}
