package main

import (
	"github.com/trasa/watchmud-message"
)

func (c *Client) handleEquipResponse(r *message.EquipResponse) {
	if r.GetSuccess() {
		UIPrintln("Ok.")
		return
	}
	switch r.GetResultCode() {
	case "NO_SLOT_GIVEN":
		UIPrintln("Equip where?")
	case "NO_TARGET":
		UIPrintln("What do you want to equip?")
	case "TARGET_NOT_FOUND":
		UIPrintln("You don't have one of those.")
	default:
		UIPrintError(r, r.GetResultCode())
	}
}
