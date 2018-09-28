package main

import (
	"github.com/trasa/watchmud-message"
)

func (c *Client) handleWearResponse(r *message.WearResponse) {
	if r.GetSuccess() {
		UIPrintln("Ok.")
		return
	}
	switch r.GetResultCode() {
	case "CANT_WEAR_THAT":
		UIPrintln("You can't wear that!")
	case "IN_USE":
		UIPrintln("You're already wearing something there.")
	case "TARGET_NOT_FOUND":
		UIPrintln("Wear what?")
	default:
		UIPrintError(r, r.GetResultCode())
	}
}
