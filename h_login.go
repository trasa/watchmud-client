package main

import (
	"errors"
	"github.com/trasa/watchmud-message"
)

func (c *Client) handleLoginResponse(resp *message.LoginResponse) error {
	if !resp.GetSuccess() {
		UIPrintln("Login Attempt Failed! ", resp.GetResultCode())
		return errors.New(resp.GetResultCode())
	} else {
		UIPrintln("Login Successful. Player name is", resp.PlayerName)
		c.playerName = resp.PlayerName

		// get the room we start off in (look request)
		msg, err := message.NewGameMessage(message.LookRequest{})
		if err != nil {
			UIPrintf("Error while trying to create Look Request after logging in (this is bad): %v\n", err)
			return err
		} else {
			c.SendMessage(msg)
		}
	}
	return nil
}
