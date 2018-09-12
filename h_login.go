package main

import (
	"errors"
	"fmt"
	"github.com/trasa/watchmud-message"
)

func (c *Client) handleLoginResponse(resp *message.LoginResponse) error {
	if !resp.GetSuccess() {
		fmt.Println("Login Attempt Failed! ", resp.GetResultCode())
		return errors.New(resp.GetResultCode())
	} else {
		fmt.Println("Login Successful. Player name is", resp.PlayerName)
		c.playerName = resp.PlayerName

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
