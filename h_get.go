package main

import (
	"fmt"
	"github.com/trasa/watchmud/message"
)

func (c *Client) handleGetResponse(r *message.GetResponse) {
	if r.IsSuccessful() {
		fmt.Println("Ok.")
		return
	}
	// otherwise...
	switch r.GetResultCode() {
	case "NO_TARGET":
		fmt.Println("Get what?")

	case "TARGET_NOT_FOUND":
		fmt.Println("You don't see one of those.")

	default:
		c.printError(r)
	}
}
