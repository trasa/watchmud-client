package main

import (
	"github.com/trasa/watchmud-message"
)

func (c *Client) handleWhoResponse(resp *message.WhoResponse) {
	if resp.GetSuccess() {
		UIPrintln("-- Who Is Here --")
		if len(resp.PlayerInfo) == 0 {
			UIPrintln("Nobody!")
		} else {
			for _, p := range resp.PlayerInfo {
				UIPrintf("%s - %s - %s\n", p.PlayerName, p.ZoneName, p.RoomName)
			}
		}
	} else {
		UIPrintError(resp, resp.GetResultCode())
	}
}
