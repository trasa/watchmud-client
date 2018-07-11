package main

import (
	"fmt"
	"github.com/trasa/watchmud-message"
)

func (c *Client) handleLoginResponse(resp *message.LoginResponse) {
	if !resp.GetSuccess() {
		fmt.Println("Login Attempt Failed! ", resp.GetResultCode())
	} else {
		fmt.Println("Login Successful. Player name is", resp.PlayerName)
		c.playerName = resp.PlayerName

		// get the room we start off in (look request)
		msg, err := message.NewGameMessage(message.LookRequest{})
		if err != nil {
			fmt.Printf("Error while trying to create Look Request after logging in (this is bad): %v\n", err)
		} else {
			c.SendMessage(msg)
		}
	}
}
