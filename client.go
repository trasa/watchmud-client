package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/trasa/watchmud/message"
	"log"
)

type Client struct {
	conn   *websocket.Conn
	quit   chan interface{}
	source chan interface{} // sends up to server
}

func Connect(serverAddress string) (*Client, error) {
	log.Printf("Connecting to %s", serverAddress)
	conn, _, err := websocket.DefaultDialer.Dial(serverAddress, nil)
	if err != nil {
		return nil, err
	}
	return NewClient(conn), nil
}

func NewClient(conn *websocket.Conn) *Client {
	c := Client{
		conn:   conn,
		quit:   make(chan interface{}),
		source: make(chan interface{}, 256),
	}
	go c.writePump()
	go c.readPump()
	return &c
}

func (c *Client) SendRequest(request interface{}) {
	requestEnvelope := message.RequestEnvelope{
		Format: "request",
		Value:  request,
	}
	/*
		j, _ := json.Marshal(requestEnv)
		log.Printf("sending \n%s\n", j)
	*/
	c.source <- requestEnvelope
}

func (c *Client) writePump() {
	defer c.conn.Close()
	for {
		select {
		case msg := <-c.source:
			if err := c.conn.WriteJSON(msg); err != nil {
				log.Println("Write Error: ", err)
				return
			}

		case <-c.quit:
			log.Println("QUIT channel message received")
			return
		}
	}
}

func (c *Client) readPump() {
	defer c.conn.Close()
	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("read error:", err)
			return
		}
		// TODO write stuff here..
		log.Printf("received: %s", msg)
		// TODO turn received message into something ...
		loginResp := &message.LoginResponse{}
		if err := json.Unmarshal(msg, loginResp); err != nil {
			log.Println("Unmarshall error: ", err)
		}
		log.Printf("loginResp %s %s", loginResp.Successful, loginResp.Player.Name)
	}
}
