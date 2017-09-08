package main

import (
	"fmt"
	"github.com/trasa/watchmud/direction"
	"github.com/trasa/watchmud/message"
)

func (c *Client) handleEnterRoomNotification(note *message.EnterRoomNotification) {
	if !note.IsSuccessful() {
		c.printError(note)
	} else {
		fmt.Println(note.PlayerName, "enters.")
	}
}

func (c *Client) handleLeaveRoomNotification(note *message.LeaveRoomNotification) {
	if !note.IsSuccessful() {
		c.printError(note)
	} else {
		dirName, err := direction.DirectionToString(note.Direction)
		if err != nil {
			fmt.Println("Error figuring out direction string for", note.Direction)
		}
		fmt.Println(note.PlayerName, "departs", dirName)
	}
}
