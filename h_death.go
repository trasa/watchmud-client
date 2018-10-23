package main

import (
	"github.com/trasa/watchmud-message"
)

func (c *Client) handleDeathNotification(resp *message.DeathNotification) {
	if resp.IsPlayer {
		if c.IsTargetYou(resp.Target) {
			// hey that's us
			UIPrintln("YOU'RE DEAD!")
			// TODO what else here ...
		} else {
			UIPrintf("%s is dead.\n", resp.Target)
		}
	} else {
		UIPrintf("The %s is dead.\n", resp.Target)
	}
}
