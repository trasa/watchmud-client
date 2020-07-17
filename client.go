package main

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/trasa/watchmud-message"
	"google.golang.org/grpc"
	"os"
)

type Client struct {
	stream      message.MudComm_SendReceiveClient
	quit        chan interface{}
	quitSignal  chan os.Signal
	source      chan *message.GameMessage // sends up to server
	isClosed    bool
	clientState *ClientState
}

func Connect(serverAddress string, port int) (*Client, error) {
	addr := fmt.Sprintf("%s:%d", serverAddress, port)
	log.Info().Msgf("Connecting to %s", addr)

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
	log.Info().Msg("Connected Successfully")
	// this starts up the writePump and the readPump
	return NewClient(stream), nil
}

// Create a new Client instance for this connection, and
// establish the writePump and readPump for that Client.
func NewClient(stream message.MudComm_SendReceiveClient) *Client {
	c := Client{
		stream:      stream,
		quit:        make(chan interface{}),
		quitSignal:  make(chan os.Signal),
		source:      make(chan *message.GameMessage, 2),
		clientState: NewClientState(),
	}
	go c.writePump()
	go c.readPump()
	return &c
}

func (c *Client) initialize() {
	// send a request for data: races
	if err := c.sendDataRequest(); err != nil {
		log.Fatal().Err(err).Msg("Failed to send data request")
		return
	}

	// await response for data: races
	// on success that handler sets clientState.inputHandler = initialInputHandler
}

func (c *Client) processInput(buf string) {
	// based on the state of the app ...
	tokens := message.Tokenize(buf)
	if len(tokens) == 0 {
		// do nothing
		return
	}

	// hand off to correct inputHandler
	c.clientState.inputHandler(c, tokens)
}

func (c *Client) sendTokens(tokens []string) {
	if c.isClosed {
		log.Warn().Msg("not sending, c.isClosed")
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
	log.Printf("Send: %T - %v", msg.Inner, msg.Inner)
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
			log.Info().Msgf("writePump: QUIT channel message received: %v", quitmessage)
			c.isClosed = true
			return
		case quitsig := <-c.quitSignal:
			log.Info().Msgf("writePump: QuitSignal received: %s", quitsig.String())
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
			log.Fatal().Err(err).Msg("failed to receive")
			return
		}
		if err := c.handleIncomingMessage(in); err != nil {
			log.Fatal().Err(err).Msg("Error handling incoming message")
			return
		}
	}
}
