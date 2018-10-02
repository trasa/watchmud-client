package main

import (
	"github.com/trasa/watchmud-message"
)

func (c *Client) handleGetResponse(r *message.GetResponse) {
	if r.Success {
		UIPrintln("Ok.")
		return
	}
	// otherwise...
	switch r.GetResultCode() {
	case "NO_TAKE":
		UIPrintln("You can't take that!")

	case "NO_TARGET":
		UIPrintln("Get what?")

	case "TARGET_NOT_FOUND":
		UIPrintln("You don't see one of those.")

	default:
		UIPrintResponseError(r, r.GetResultCode())
	}
}

func (c *Client) handleGetNotification(n *message.GetNotification) {
	if n.Success {
		// TODO clauses, articles, plural and so on...
		UIPrintf("%s gets a %s.\n", n.PlayerName, n.Target)
		return
	}
	// weird error case
	UIPrintResponseError(n, n.GetResultCode())
}
