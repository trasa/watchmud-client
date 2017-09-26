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

func (c *Client) handleDropNotification(n *message.DropNotification) {
	if n.IsSuccessful() {
		// TODO clauses, articles, plural and so on...
		fmt.Printf("%s drops a %s.\n", n.PlayerName, n.Target)
		return
	}
	// weird error case
	c.printError(n)
}
