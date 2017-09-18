package main

import (
	"fmt"
	"github.com/trasa/watchmud/direction"
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
		fmt.Println(note.PlayerName, "departs", dirName+".")
	}
}
