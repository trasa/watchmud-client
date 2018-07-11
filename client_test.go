package main

import "github.com/trasa/watchmud-message"

func NewTestClient() *Client {
	return &Client{
		source: make(chan *message.GameMessage, 2),
	}
}
