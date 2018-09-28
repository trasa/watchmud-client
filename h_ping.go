package main

import (
	"github.com/trasa/watchmud-message"
)

func (c *Client) handlePing(ping *message.Ping) {
	UIPrintf("Ping! %s", ping.Target)
	// TODO send pong
}

func (c *Client) handlePong(pong *message.Pong) {
	UIPrintf("Pong! %s", pong.Target)
}
