package main

import (
	"fmt"
	"github.com/trasa/watchmud-message"
)

func (c *Client) handleWearResponse(r *message.WearResponse) {
	if r.GetSuccess() {
		fmt.Println("Ok.")
		return
	}
	switch r.GetResultCode() {
	case "CANT_WEAR_THAT":
		fmt.Println("You can't wear that!")
	case "IN_USE":
		fmt.Println("You're already wearing something there.")
	case "TARGET_NOT_FOUND":
		fmt.Println("Wear what?")
	default:
		c.printError(r, r.GetResultCode())
	}
}
