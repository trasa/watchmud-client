package main

import (
	"fmt"
	"github.com/trasa/watchmud/message"
)

func (c *Client) handleLoginResponse(resp *message.LoginResponse) {
	if !resp.IsSuccessful() {
		fmt.Println("Login Attempt Failed! ", resp.GetResultCode())
	} else {
		fmt.Println("Login Successful. Player name is", resp.Player.Name)
		c.playerData = resp.Player
	}
}
