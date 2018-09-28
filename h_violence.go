package main

import (
	"github.com/trasa/watchmud-message"
)

func (c *Client) handleViolenceNotification(r *message.ViolenceNotification) {

	fighterIsYou := c.IsTargetYou(r.Fighter)
	fighteeIsYou := c.IsTargetYou(r.Fightee)

	if r.SuccessfulHit {
		if fighterIsYou {
			UIPrintf("You hit %s! (%d damage)\n", r.Fightee, r.Damage)
		} else if fighteeIsYou {
			UIPrintf("%s hits you! (%d damage)\n", r.Fighter, r.Damage)
		} else {
			UIPrintf("%s hits %s for %d damage.\n", r.Fighter, r.Fightee, r.Damage)
		}
	} else {
		if fighterIsYou {
			UIPrintf("You try to hit %s, but miss.\n", r.Fightee)
		} else if fighteeIsYou {
			UIPrintf("%s tries to hit you, but fails.\n", r.Fighter)
		} else {
			UIPrintf("%s tries to hit %s but misses.\n", r.Fighter, r.Fightee)
		}
	}
}
