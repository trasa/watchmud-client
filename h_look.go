package main

import (
	"github.com/trasa/watchmud/message"
)

func (c *Client) handleLookResponse(resp *message.LookResponse) {
	if !resp.GetSuccess() {
		c.printError(resp, resp.GetResultCode())
	} else {
		c.printRoom(resp.GetRoomDescription())
	}
}
