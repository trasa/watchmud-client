package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/trasa/watchmud/message"
	"log"
	"os"
	"time"
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

func (c *Client) SendLine(line string) {
	requestEnvelope := message.RequestEnvelope{
		Format: "line",
		Value:  line,
	}
	c.source <- requestEnvelope
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

// Read stdin for line input and send to the server
// until we receive a command like /q, in which case
// this terminates.
func (c *Client) readStdin() {
	scanner := bufio.NewScanner(os.Stdin)
	c.printPrompt()
	for scanner.Scan() {
		line := scanner.Text()
		if line == "/q" {
			c.quit <- "exit now!"
			time.Sleep(1 * time.Second)
			break
		}
		c.SendLine(line)
		c.printPrompt()
	}
}

func (c *Client) printPrompt() {
	fmt.Print("> ")
}

func (c *Client) handleIncomingResponse(resp message.Response) {
	switch resp.(type) {
	case *message.LoginResponse:
		c.handleLoginResponse(resp.(*message.LoginResponse))

	default:
		log.Println("unknown response type", resp)
	}
}
