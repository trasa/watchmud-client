package main

import (
	"github.com/gorilla/websocket"

	"encoding/json"
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

	// connect to websocket
	log.Printf("Connecting to %s", SERVER_ADDR)
	conn, _, err := websocket.DefaultDialer.Dial(SERVER_ADDR, nil)
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}
	defer conn.Close()
	done := make(chan struct{})

	go func() {
		defer conn.Close()
		defer close(done)
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("read error: ", err)
				return
			}
			log.Printf("recv: %s", msg)
			loginResp := &message.LoginResponse{}
			if err := json.Unmarshal(msg, loginResp); err != nil {
				log.Println("Unmarshall error: ", err)
			}
			log.Printf("loginResp %s %s", loginResp.Successful, loginResp.Player.Name)
		}
	}()

	// send login request
	playerName := "somedood"
	password := "NotImplemented"

	loginReq := message.LoginRequest{
		Request:    message.RequestBase{MessageType: "login"},
		PlayerName: playerName,
		Password:   password,
	}
	requestEnv := message.RequestEnvelope{
		Format: "request",
		Value:  loginReq,
	}
	j, _ := json.Marshal(requestEnv)
	log.Printf("sending \n%s\n", j)

	conn.WriteJSON(requestEnv)
	time.Sleep(time.Second * 10)
}
