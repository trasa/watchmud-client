package main

import (
	"github.com/gorilla/websocket"
	"github.com/trasa/watchmud/message"
	"log"
)

type Client struct {
	conn       *websocket.Conn
	quit       chan interface{}
	source     chan interface{}   // sends up to server
	playerData message.PlayerData // who am I anyway
}

func Connect(serverAddress string) (*Client, error) {
	log.Printf("Connecting to %s", serverAddress)
	conn, _, err := websocket.DefaultDialer.Dial(serverAddress, nil)
	if err != nil {
		return nil, err
	}
	// this starts up the writePump and the readPump
	return NewClient(conn), nil
}

// Create a new Client instance for this connection, and
// establish the writePump and readPump for that Client.
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
		log.Printf("raw received: %s", msg)
		if r, err := message.TranslateToResponse(msg); err != nil {
			log.Println("unmarshal / translate error", err)
		} else {
			c.handleIncomingResponse(r)
		}
	}
}

func (c *Client) handleIncomingResponse(resp message.Response) {
	switch resp.(type) {
	case *message.LoginResponse:
		c.handleLoginResponse(resp.(*message.LoginResponse))

	default:
		log.Println("unknown response type", resp)
	}
}
