package main

import (
	"github.com/trasa/watchmud/message"
)

func (c *Client) handleLookResponse(resp *message.LookResponse) {
	if !resp.IsSuccessful() {
		c.printError(resp)
	} else {
		c.printRoom(resp.RoomDescription)
	}
}
