package main

import (
	"fmt"
	"github.com/trasa/watchmud/direction"
	"github.com/trasa/watchmud/message"
)

func (c *Client) handleLookResponse(resp *message.LookResponse) {
	if !resp.IsSuccessful() {
		fmt.Println("Error Looking:", resp.GetResultCode())
	} else {
		fmt.Println(resp.RoomDescription.Name)
		fmt.Println()
		fmt.Println(resp.RoomDescription.Description)
		// other players
		if len(resp.RoomDescription.Players) > 0 {
			fmt.Println("Other Players:")
			for _, p := range resp.RoomDescription.Players {
				fmt.Println(p)
			}
		}
		// obvious exits
		if exits, err := direction.ExitsToString(resp.RoomDescription.Exits); err == nil {
			fmt.Println("Obvious Exits:", exits)
		} else {
			fmt.Println("Error Getting exits:", err)
		}
	}
}
