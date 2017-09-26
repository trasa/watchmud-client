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

func (c *Client) handleGetNotification(n *message.GetNotification) {
	if n.IsSuccessful() {
		// TODO clauses, articles, plural and so on...
		fmt.Printf("%s gets a %s.\n", n.PlayerName, n.Target)
		return
	}
	// weird error case
	c.printError(n)
}
