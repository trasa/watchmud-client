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
