package main

import (
	"fmt"
	"github.com/trasa/watchmud/message"
)

func (c *Client) handleDropResponse(r *message.DropResponse) {
	if r.IsSuccessful() {
		fmt.Println("Dropped.")
		return
	}
	//otherwise...
	switch r.GetResultCode() {
	case "NO_TARGET":
		fmt.Println("Drop what?")

	case "TARGET_NOT_FOUND":
		fmt.Println("You don't have one of those.")

	default:
		c.printError(r)
	}
}
