package main

import message "github.com/trasa/watchmud-message"

func (c *Client) handleRestoreResponse(r *message.RestoreResponse) {
	// TODO what if you don't have permissions
	if r.GetSuccess() {
		UIPrintln("Done.")
		return
	}
	UIPrintResponseError(r, r.GetResultCode())
}

func (c *Client) handleRestoreNotification(r *message.RestoreNotification) {
	UIPrintf("%s looks much healthier.\n", r.GetTarget())
}
