package main

import (
	"fmt"
	"github.com/trasa/watchmud-message"
)

func (c *Client) handleWhoResponse(resp *message.WhoResponse) {
	if resp.GetSuccess() {
		fmt.Println("-- Who Is Here --")
		if len(resp.PlayerInfo) == 0 {
			fmt.Println("Nobody!")
		} else {
			for _, p := range resp.PlayerInfo {
				fmt.Printf("%s - %s - %s\n", p.PlayerName, p.ZoneName, p.RoomName)
			}
		}
	} else {
		c.printError(resp, resp.GetResultCode())
	}
}
