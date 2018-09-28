package main

import (
	"github.com/trasa/watchmud-message"
	"github.com/trasa/watchmud-message/direction"
)

func (c *Client) handleExitsResponse(r *message.ExitsResponse) {
	if !r.GetSuccess() {
		UIPrintError(r, r.GetResultCode())
	} else {
		UIPrintln("Obvious Exits:")
		if len(r.ExitInfo) == 0 {
			UIPrintln("None!")
		} else {
			for _, rexit := range r.ExitInfo {
				if dirName, err := direction.DirectionToString(direction.Direction(rexit.Direction)); err == nil {
					UIPrintf("%s - %s\n", dirName, rexit.RoomName)
				} else {
					UIPrintf("Error with direction name: %d %s - %s\n", rexit.Direction, rexit.RoomName, err)
				}
			}
		}
	}
}
