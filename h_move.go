package main

import (
	"fmt"
	"github.com/trasa/watchmud/message"
)

func (c *Client) handleMoveResponse(resp *message.MoveResponse) {
	if resp.IsSuccessful() {
		c.printRoom(resp.RoomDescription)
	} else if resp.GetResultCode() == "CANT_GO_THAT_WAY" {
		fmt.Println("There's no exit that way.")
	} else {
		c.printError(resp)
	}
}
