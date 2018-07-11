package main

import (
	"fmt"
	"github.com/trasa/watchmud-message"
)

func (c *Client) handlePing(ping *message.Ping) {
	fmt.Printf("Ping! %s", ping.Target)
	// TODO send pong
}

func (c *Client) handlePong(pong *message.Pong) {
	fmt.Printf("Pong! %s", pong.Target)
}
