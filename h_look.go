package main

import (
	"fmt"
	"github.com/trasa/watchmud/message"
)

func (c *Client) handleLookResponse(resp *message.LookResponse) {
	if !resp.IsSuccessful() {
		fmt.Println("Error Looking:", resp.GetResultCode())
	} else {
		c.printRoom(resp.RoomDescription)

	}
}
