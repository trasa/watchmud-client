package main

import message "github.com/trasa/watchmud-message"

func (c *Client) handleLoadResponse(r *message.LoadResponse) {
	if r.Success {
		UIPrintln("Ok.")
	} else {
		UIPrintf("Load Request Failed: %s", r.ResultCode)
	}
}
