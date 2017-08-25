package main

import (
	"github.com/trasa/watchmud/message"
	"log"
	"os"
	"os/signal"
	"time"
)

const SERVER_ADDR = "ws://localhost:8888/ws"

func main() {
	// TODO read from yaml configuration or something
	// TODO override with command line args

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// connect client
	client, err := Connect(SERVER_ADDR)
	if err != nil {
		log.Fatal("Failed to connect", err)
	}

	// send login request
	playerName := "somedood"
	password := "NotImplemented"

	loginReq := message.LoginRequest{
		Request:    message.RequestBase{MessageType: "login"},
		PlayerName: playerName,
		Password:   password,
	}
	client.SendRequest(loginReq)
	// TODO do something more interesting than sleep and exit...
	time.Sleep(time.Second * 10)
}
