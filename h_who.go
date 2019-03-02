package main

import (
	"fmt"
	"github.com/trasa/watchmud-message"
	"strings"
)

func (c *Client) handleWhoResponse(resp *message.WhoResponse) {
	if resp.GetSuccess() {
		var str strings.Builder
		str.WriteString("-- Who Is Here --\n")
		if len(resp.PlayerInfo) == 0 {
			str.WriteString("Nobody!\n")
		} else {
			for _, p := range resp.PlayerInfo {
				str.WriteString(fmt.Sprintf("%s - %s - %s\n", p.PlayerName, p.ZoneName, p.RoomName))
			}
		}
		UIPrint(str)
	} else {
		UIPrintResponseError(resp, resp.GetResultCode())
	}
}
