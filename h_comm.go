package main

import (
	"fmt"
	"github.com/trasa/watchmud/message"
)

func (c *Client) handleSayResponse(resp *message.SayResponse) {
	if resp.IsSuccessful() {
		fmt.Printf("You say '%s'\n", resp.Value)
	} else {
		if resp.GetResultCode() == "NOT_IN_A_ROOM" {
			fmt.Println("You yell into the darkness.")
		} else {
			c.printError(resp)
		}
	}
}

func (c *Client) handleSayNotification(note *message.SayNotification) {
	if note.IsSuccessful() {
		fmt.Printf("%s says '%s'.\n", note.Sender, note.Value)
	} else {
		c.printError(note)
	}
}

func (c *Client) handleTellNotification(note *message.TellNotification) {
	if note.IsSuccessful() {
		fmt.Printf("%s tells you '%s'.\n", note.Sender, note.Value)
	} else {
		c.printError(note)
	}
}

func (c *Client) handleTellResponse(resp *message.TellResponse) {
	if resp.IsSuccessful() {
		fmt.Println("sent.")
	} else if resp.GetResultCode() == "TO_PLAYER_NOT_FOUND" {
		fmt.Println("Nobody here by that name.")
	} else {
		c.printError(resp)
	}
}

func (c *Client) handleTellAllResponse(resp *message.TellAllResponse) {
	if resp.IsSuccessful() {
		fmt.Println("sent.")
	} else {
		c.printError(resp)
	}
}

func (c *Client) handleTellAllNotification(note *message.TellAllNotification) {
	if note.IsSuccessful() {
		fmt.Printf("tell_all %s> %s", note.Sender, note.Value)
	} else {
		c.printError(note)
	}
}
