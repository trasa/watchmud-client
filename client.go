package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/trasa/watchmud-message"
	"google.golang.org/grpc"
	"log"
	"os"
)

type Client struct {
	stream     message.MudComm_SendReceiveClient
	quit       chan interface{}
	quitSignal chan os.Signal
	source     chan *message.GameMessage // sends up to server
	isClosed   bool
	playerName string
}

func Connect(serverAddress string, port int) (*Client, error) {
	addr := fmt.Sprintf("%s:%d", serverAddress, port)
	log.Printf("Connecting to %s", addr)

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		return nil, err
	}
	rpcClient := message.NewMudCommClient(conn)
	stream, err := rpcClient.SendReceive(context.Background())
	if err != nil {
		return nil, err
	}
	// this starts up the writePump and the readPump
	return NewClient(stream), nil
}

// Create a new Client instance for this connection, and
// establish the writePump and readPump for that Client.
func NewClient(stream message.MudComm_SendReceiveClient) *Client {
	c := Client{
		stream:     stream,
		quit:       make(chan interface{}),
		quitSignal: make(chan os.Signal),
		source:     make(chan *message.GameMessage, 2),
	}
	go c.writePump()
	go c.readPump()
	return &c
}

func (c *Client) sendTokens(tokens []string) {
	if c.isClosed {
		log.Println("not sending, c.isClosed")
		return
	}
	// translate line into message
	msg, err := message.TranslateLineToMessage(tokens)
	if err != nil {
		// TODO better error message
		log.Printf("Error creating New Game Message for payload: %v", err)
	} else {
		c.source <- msg
	}
}

func (c *Client) SendMessage(msg *message.GameMessage) {
	c.source <- msg
}

func (c *Client) writePump() {
	for {
		select {
		case msg := <-c.source:
			if err := c.stream.Send(msg); err != nil {
				log.Printf("RPC Write Error: %v", err)
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
	for {
		in, err := c.stream.Recv()
		if err != nil {
			c.quit <- fmt.Sprint("read error ", err)
			c.isClosed = true
			log.Fatalf("failed to receive: %v", err)
			return
		}
		if err := c.handleIncomingMessage(in); err != nil {
			log.Fatalf("Error handling incoming message: %v", err)
			return
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
		tokens := message.Tokenize(line)
		if len(tokens) == 0 {
			// do nothing...
		} else if tokens[0] == "quit" || tokens[0] == "/q" {
			c.quit <- "QUIT command"
			return // <- exits the loop and the app
		} else if tokens[0] == "help" {
			printHelp(tokens)
		} else {
			// send to server
			c.sendTokens(tokens)
		}
	}
}
