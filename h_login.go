package main

import (
	"errors"
	"github.com/trasa/watchmud-message"
)

func (c *Client) doLogin(tokens []string) error {
	// send login request
	password := "NotImplemented"
	var playerName string
	if len(tokens) < 2 {
		playerName = "somedood"
	} else {
		playerName = tokens[1]
	}

	loginReq := message.LoginRequest{
		PlayerName: playerName,
		Password:   password,
	}
	loginMsg, err := message.NewGameMessage(loginReq)
	if err != nil {
		return err
	}
	c.SendMessage(loginMsg)
	return nil
}

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
