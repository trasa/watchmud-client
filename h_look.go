package main

import (
	"fmt"
	"github.com/trasa/watchmud/message"
	"github.com/trasa/watchmud/direction"
)

func (c *Client) handleLookResponse(resp *message.LookResponse) {
	if !resp.IsSuccessful() {
		fmt.Println("Error Looking:", resp.GetResultCode())
	} else {
		fmt.Println(resp.Name)
		fmt.Println("\n", resp.Description)
		// other players
		if len(resp.Players) > 0 {
			fmt.Println("Other Players:")
			for _, p := range resp.Players {
				fmt.Println(p)
			}
		}
		// obvious exits
		if exits, err := direction.ExitsToString(resp.Exits); err == nil {
			fmt.Println("Obvious Exits:", exits)
		} else {
			fmt.Println("Error Getting exits:", err)
		}
	}
}
