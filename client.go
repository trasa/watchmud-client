package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/trasa/watchmud/message"
	"log"
	"os"
)

type Client struct {
	conn       *websocket.Conn
	quit       chan interface{}
	quitSignal chan os.Signal
	source     chan interface{} // sends up to server
	isClosed   bool
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
		conn:       conn,
		quit:       make(chan interface{}),
		quitSignal: make(chan os.Signal),
		source:     make(chan interface{}, 2),
	}
	go c.writePump()
	go c.readPump()
	return &c
}

func (c *Client) SendLine(line string) {
	if c.isClosed {
		log.Println("not sending, c.isClosed", line)
		return
	}
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
	//noinspection GoBoolExpressions
	if COMM_LOGGING_ENABLED {
		j, _ := json.Marshal(requestEnvelope)
		log.Printf("sending \n%s\n", j)
	}
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

		case quitmessage := <-c.quit:
			log.Println("writePump: QUIT channel message received:", quitmessage)
			c.isClosed = true
			return
		case quitsig := <-c.quitSignal:
			log.Println("writePump: QuitSignal received:", quitsig.String())
			c.isClosed = true
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
			c.quit <- fmt.Sprint("read error ", err)
			c.isClosed = true
			return
		}
		//noinspection GoBoolExpressions
		if COMM_LOGGING_ENABLED {
			log.Printf("raw received: %s", msg)
		}
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
	// seriously annoying: this blocks forever, even if SIGINT
	// has been sent. There doesn't seem to be a way to set up a
	// signal handler here to break us out of this for loop if
	// received
	for scanner.Scan() {
		if c.isClosed {
			log.Println("c.isClosed")
			return
		}
		line := scanner.Text()
		if line == "/q" {
			c.quit <- "QUIT command"
			return
		}
		c.SendLine(line)
	}
}
