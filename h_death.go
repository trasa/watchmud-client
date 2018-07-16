package main

import (
	"fmt"
	"github.com/trasa/watchmud-message"
)

func (c *Client) handleDeathNotification(resp *message.DeathNotification) {
	if resp.IsPlayer {
		if resp.Target == c.playerName {
			// hey that's us
			fmt.Println("YOU'RE DEAD!")
			// TODO what else here ...
		} else {
			fmt.Printf("%s is dead.\n", resp.Target)
		}
	} else {
		fmt.Printf("The %s is dead.\n", resp.Target)
	}
}
