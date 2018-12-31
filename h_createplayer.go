package main

import (
	"errors"
	"fmt"
	"github.com/trasa/watchmud-message"
)

func (c *Client) sendCreatePlayerRequest() error {
	password := "NotImplemented"
	// TODO implement races and classes...
	createReq := message.CreatePlayerRequest{
		PlayerName: c.clientState.playerName,
		Password:   password,
		Race:       c.clientState.race,
		Class:      c.clientState.class,
	}
	createMsg, err := message.NewGameMessage(createReq)
	if err != nil {
		return err
	}
	c.SendMessage(createMsg)
	return nil
}

func (c *Client) handleCreatePlayerResponse(resp *message.CreatePlayerResponse) error {
	// TODO deal with case where name chosen has already been taken -
	// for that matter, most errors should do this:
	// don't lose all the player's choices, allow them to fix up their Request,
	// and then try to re-send it.
	if !resp.GetSuccess() {
		UIPrintln("Create Player Attempt Failed! ", resp.GetResultCode())
		return errors.New(resp.GetResultCode())
		// TODO recover from create error gracefully
	} else {
		UIPrintln("Create Player successful. Player name is", resp.PlayerName)
		c.clientState.inputHandler = gameInputHandler
		// get the room we start off in (look request)
		msg, err := message.NewGameMessage(message.LookRequest{})
		if err != nil {
			fmt.Printf("Error while trying to create Look Request after logging in (this is bad): %v\n", err)
			return err
		} else {
			c.SendMessage(msg)
		}
	}
	return nil
}
