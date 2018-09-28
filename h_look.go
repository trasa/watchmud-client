package main

import (
	"github.com/trasa/watchmud-message"
)

func (c *Client) handleLookResponse(resp *message.LookResponse) {
	if !resp.GetSuccess() {
		UIPrintError(resp, resp.GetResultCode())
	} else {
		UIPrintRoom(resp.GetRoomDescription())
	}
}
