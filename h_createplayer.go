package main

import (
	"errors"
	"github.com/trasa/watchmud-message"
)

func (c *Client) handleCreatePlayerResponse(resp *message.CreatePlayerResponse) error {
	if !resp.GetSuccess() {
		UIPrintln("Create Player Attempt Failed! ", resp.GetResultCode())
		return errors.New(resp.GetResultCode())
	} else {
		UIPrintln("Create Player successful. Player name is", resp.PlayerName)

		// TODO clientState
		// TODO whatever else we need to do on creating a new user...

	}
	return nil
}
