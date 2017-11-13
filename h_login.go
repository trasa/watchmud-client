package main

import (
	"fmt"
	"github.com/trasa/watchmud/message"
)

func (c *Client) handleLoginResponse(resp *message.LoginResponse) {
	if !resp.GetSuccess() {
		fmt.Println("Login Attempt Failed! ", resp.GetResultCode())
	} else {
		fmt.Println("Login Successful. Player name is", resp.PlayerName)
		c.playerName = resp.PlayerName
		// get the room we start off in
		c.SendLine("look")
	}
}
