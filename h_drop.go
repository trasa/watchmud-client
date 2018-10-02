package main

import (
	"github.com/trasa/watchmud-message"
)

func (c *Client) handleDropResponse(r *message.DropResponse) {
	if r.GetSuccess() {
		UIPrintln("Dropped.")
		return
	}
	//otherwise...
	switch r.GetResultCode() {
	case "NO_TARGET":
		UIPrintln("Drop what?")

	case "TARGET_NOT_FOUND":
		UIPrintln("You don't have one of those.")

	default:
		UIPrintResponseError(r, r.GetResultCode())
	}
}

func (c *Client) handleDropNotification(n *message.DropNotification) {
	if n.GetSuccess() {
		// TODO clauses, articles, plural and so on...
		UIPrintf("%s drops a %s.\n", n.PlayerName, n.Target)
		return
	}
	// weird error case
	UIPrintResponseError(n, n.GetResultCode())
}
