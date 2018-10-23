package main

import (
	"fmt"
	"github.com/trasa/watchmud-message"
)

func (c *Client) doLogin(tokens []string) error {
	if c.clientState.currentState != Initial {
		UIPrintf("You cannot 'login', you are in state '%s'\n", c.clientState.currentState)
		return nil
	}
	c.clientState.currentState = Login

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
		// we are not logged in, don't return error as this will terminate the client.
		c.clientState.handleLoginFailed()
		return nil
	} else {
		UIPrintln("Login Successful. Player name is", resp.PlayerName)
		// we are logged in
		c.clientState.handleLoginSucceeded(resp.PlayerName)

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
