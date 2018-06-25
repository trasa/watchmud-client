package main

import (
	"fmt"
	"github.com/trasa/watchmud/direction"
	"github.com/trasa/watchmud/message"
)

func (c *Client) handleExitsResponse(r *message.ExitsResponse) {
	if !r.GetSuccess() {
		c.printError(r, r.GetResultCode())
	} else {
		fmt.Println("Obvious Exits:")
		if len(r.ExitInfo) == 0 {
			fmt.Println("None!")
		} else {
			for _, rexit := range r.ExitInfo {
				if dirName, err := direction.DirectionToString(direction.Direction(rexit.Direction)); err == nil {
					fmt.Printf("%s - %s\n", dirName, rexit.RoomName)
				} else {
					fmt.Printf("Error with direction name: %d %s - %s\n", rexit.Direction, rexit.RoomName, err)
				}
			}
		}
	}
}
