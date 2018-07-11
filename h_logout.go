package main

import (
	"fmt"
	"github.com/trasa/watchmud-message"
)

func (c *Client) handleLogoutNotification(resp *message.LogoutNotification) {
	var name string
	if resp.PlayerName == "" {
		name = "Someone"
	} else {
		name = resp.PlayerName
	}
	fmt.Printf("%s disappears.\n", name)
}
