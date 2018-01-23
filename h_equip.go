package main

import (
	"fmt"
	"github.com/trasa/watchmud/message"
)

func (c *Client) handleEquipResponse(r *message.EquipResponse) {
	if r.GetSuccess() {
		fmt.Println("Ok.")
		return
	}
	switch r.GetResultCode() {
	case "NO_SLOT_GIVEN":
		fmt.Println("Equip where?")
	case "NO_TARGET":
		fmt.Println("What do you want to equip?")
	case "TARGET_NOT_FOUND":
		fmt.Println("You don't have one of those.")
	default:
		c.printError(r, r.GetResultCode())
	}
}
