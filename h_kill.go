package main

import (
	"github.com/trasa/watchmud-message"
)

func (c *Client) handleKillResponse(r *message.KillResponse) {
	if r.Success {
		UIPrintln("Ok.")
		return
	}
	switch r.GetResultCode() {
	case "ALREADY_FIGHTING":
		UIPrintln("You're already in a fight!")
	case "TARGET_NOT_FOUND":
		UIPrintln("You don't see them here.")
	case "NO_FIGHT_ROOM":
		UIPrintln("You feel a sense of peace.")
	case "NO_FIGHT":
		UIPrintln("Something about them convinces you not to fight.")
	default:
		UIPrintResponseError(r, r.GetResultCode())
	}

}
