package main

import "github.com/trasa/watchmud-message"

func (c *Client) handleRecallResponse(r *message.RecallResponse) {
	if r.GetSuccess() {
		UIPrintRoom(r.RoomDescription)
		return
	}
	UIPrintResponseError(r, r.GetResultCode())
}
