package main

import (
	"fmt"
	"github.com/trasa/watchmud-message"
	"github.com/trasa/watchmud-message/direction"
	"strings"
)

func (c *Client) handleExitsResponse(r *message.ExitsResponse) {
	if !r.GetSuccess() {
		UIPrintResponseError(r, r.GetResultCode())
	} else {
		var str strings.Builder
		str.WriteString("Obvious Exits:\n")
		if len(r.ExitInfo) == 0 {
			str.WriteString("None!\n")
		} else {
			for _, rexit := range r.ExitInfo {
				if dirName, err := direction.DirectionToString(direction.Direction(rexit.Direction)); err == nil {
					str.WriteString(fmt.Sprintf("%s - %s\n", dirName, rexit.RoomName))
				} else {
					str.WriteString(fmt.Sprintf("Error with direction name: %d %s - %s\n", rexit.Direction, rexit.RoomName, err))
				}
			}
		}
		UIPrint(str)
	}
}
