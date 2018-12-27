package main

import (
	"errors"
	"fmt"
	"github.com/trasa/watchmud-message"
)

func (c *Client) beginCreatePlayer(token []string) error {
	if c.clientState.currentState != Initial {
		return errors.New(fmt.Sprintf("You cannot 'begin createPlayer', you are in state '%s'\n", c.clientState.currentState))
	}
	c.clientState.currentState = CreatePlayer

	UIPrintln("What is your player's name?")
	return nil
}

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
