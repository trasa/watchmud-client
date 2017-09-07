package main

import (
	"fmt"
	"github.com/trasa/watchmud/direction"
	"github.com/trasa/watchmud/message"
)

func (c *Client) handleExitsResponse(r *message.ExitsResponse) {
	if !r.IsSuccessful() {
		fmt.Println("Error:", r.GetResultCode())
	} else {
		fmt.Println("Obvious Exits:")
		if len(r.ExitInfo) == 0 {
			fmt.Println("None!")
		} else {
			for shortDir, roomName := range r.ExitInfo {
				if dirName, err := direction.AbbreviationToString(shortDir); err == nil {
					fmt.Printf("%s - %s\n", dirName, roomName)
				} else {
					fmt.Printf("Error with direction name: %s %s - %s\n", shortDir, roomName, err)
				}
			}
		}
	}
}
