package main

import (
	"github.com/trasa/watchmud/message"
	"fmt"
)

func (c *Client) handleEnterRoomNotification(note *message.EnterRoomNotification) {
	if !note.IsSuccessful() {
		c.printError(note)
	} else {
		fmt.Println(note.PlayerName, "enters.")
	}
}