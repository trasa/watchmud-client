package main

import (
	"fmt"
	"github.com/trasa/watchmud/direction"
	"github.com/trasa/watchmud/message"
)

func (c *Client) handleMoveResponse(resp *message.MoveResponse) {
	if resp.GetSuccess() {
		c.printRoom(resp.RoomDescription)
	} else if resp.GetResultCode() == "CANT_GO_THAT_WAY" {
		fmt.Println("There's no exit that way.")
	} else {
		c.printError(resp, resp.GetResultCode())
	}
}

func (c *Client) handleEnterRoomNotification(note *message.EnterRoomNotification) {
	if !note.GetSuccess() {
		c.printError(note, note.GetResultCode())
	} else {
		fmt.Println(note.Name, "enters.")
	}
}

func (c *Client) handleLeaveRoomNotification(note *message.LeaveRoomNotification) {
	if !note.GetSuccess() {
		c.printError(note, note.GetResultCode())
	} else {
		dirName, err := direction.DirectionToString(direction.Direction(note.Direction))
		if err != nil {
			fmt.Println("Error figuring out direction string for", note.Direction)
		}
		fmt.Println(note.Name, "departs", dirName+".")
	}
}
