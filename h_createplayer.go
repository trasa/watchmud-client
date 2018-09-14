package main

import (
	"errors"
	"fmt"
	"github.com/trasa/watchmud-message"
)

func (c *Client) handleCreatePlayerResponse(resp *message.CreatePlayerResponse) error {
	if !resp.GetSuccess() {
		fmt.Println("Create Player Attempt Failed! ", resp.GetResultCode())
		return errors.New(resp.GetResultCode())
	} else {
		fmt.Println("Create Player successful. Player name is", resp.PlayerName)
		c.playerName = resp.PlayerName
		// TODO whatever else we need to do on creating a new user...

	}
	return nil
}
