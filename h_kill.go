package main

import (
	"fmt"
	"github.com/trasa/watchmud/message"
)

func (c *Client) handleKillResponse(r *message.KillResponse) {
	if r.Success {
		fmt.Println("Ok.")
		return
	}
	switch r.GetResultCode() {
	case "ALREADY_FIGHTING":
		fmt.Println("You're already in a fight!")
	case "TARGET_NOT_FOUND":
		fmt.Println("You don't see them here.")
	case "NO_FIGHT_ROOM":
		fmt.Println("You feel a sense of peace.")
	case "NO_FIGHT":
		fmt.Println("Something about them convinces you not to fight.")
	default:
		c.printError(r, r.GetResultCode())
	}

}
