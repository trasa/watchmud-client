package main

import (
	"github.com/trasa/watchmud-message"
)

func (c *Client) handleLookResponse(resp *message.LookResponse) {
	if resp.GetSuccess() {
		UIPrintRoom(resp.GetRoomDescription())
	} else {
		UIPrintResponseError(resp, resp.GetResultCode())
	}
}
