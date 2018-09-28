package main

import (
	"github.com/trasa/watchmud-message"
)

func (c *Client) handleLogoutNotification(resp *message.LogoutNotification) {
	var name string
	if resp.PlayerName == "" {
		name = "Someone"
	} else {
		name = resp.PlayerName
	}
	UIPrintf("%s disappears.\n", name)
}
