package main

import (
	"github.com/trasa/watchmud-message"
	"github.com/trasa/watchmud-message/direction"
)

func (c *Client) handleMoveResponse(resp *message.MoveResponse) {
	if resp.GetSuccess() {
		UIPrintRoom(resp.RoomDescription)
	} else if resp.GetResultCode() == "CANT_GO_THAT_WAY" {
		UIPrintln("There's no exit that way.")
	} else if resp.GetResultCode() == "IN_A_FIGHT" {
		UIPrintln("Your fighting for your life!")
	} else {
		UIPrintResponseError(resp, resp.GetResultCode())
	}
}

func (c *Client) handleEnterRoomNotification(note *message.EnterRoomNotification) {
	if !note.GetSuccess() {
		UIPrintResponseError(note, note.GetResultCode())
	} else {
		UIPrintln(note.Name, "enters.")
	}
}

func (c *Client) handleLeaveRoomNotification(note *message.LeaveRoomNotification) {
	if !note.GetSuccess() {
		UIPrintResponseError(note, note.GetResultCode())
	} else {
		dirName, err := direction.DirectionToString(direction.Direction(note.Direction))
		if err != nil {
			UIPrintln("Error figuring out direction string for", note.Direction)
		}
		UIPrintln(note.Name, "departs", dirName+".")
	}
}
