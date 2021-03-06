package main

import (
	"fmt"
	"github.com/trasa/watchmud-message"
	"strings"
)

// send a login request based on state
func (c *Client) sendLoginRequest() error {
	// send login request
	password := "NotImplemented"

	loginReq := message.LoginRequest{
		PlayerName: c.clientState.playerName,
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
		var str strings.Builder
		str.WriteString("Login Attempt Failed! " + resp.GetResultCode() + "\n")
		// we are not logged in, don't return error as this will terminate the client.
		str.WriteString("So who are you?\n")
		c.clientState.inputHandler = loginNameInputHandler
		UIPrint(str)
		return nil
	} else {

		// TODO, set client state values here

		UIPrintln("Login Successful. Player name is", resp.PlayerName)
		// we are logged in - time for normal input?
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
