package main

import (
	"github.com/trasa/watchmud-message"
)

func (c *Client) handleSayResponse(resp *message.SayResponse) {
	if resp.Success {
		UIPrintf("You say '%s'\n", resp.Value)
	} else {
		if resp.GetResultCode() == "NOT_IN_A_ROOM" {
			UIPrintln("You yell into the darkness.")
		} else {
			UIPrintResponseError(resp, resp.GetResultCode())
		}
	}
}

func (c *Client) handleSayNotification(note *message.SayNotification) {
	if note.Success {
		UIPrintf("%s says '%s'.\n", note.Sender, note.Value)
	} else {
		UIPrintResponseError(note, note.GetResultCode())
	}
}

func (c *Client) handleTellNotification(note *message.TellNotification) {
	if note.GetSuccess() {
		UIPrintf("%s tells you '%s'.\n", note.Sender, note.Value)
	} else {
		UIPrintResponseError(note, note.GetResultCode())
	}
}

func (c *Client) handleTellResponse(resp *message.TellResponse) {
	if resp.GetSuccess() {
		UIPrintln("sent.")
	} else if resp.GetResultCode() == "TO_PLAYER_NOT_FOUND" {
		UIPrintln("Nobody here by that name.")
	} else {
		UIPrintResponseError(resp, resp.GetResultCode())
	}
}

func (c *Client) handleTellAllResponse(resp *message.TellAllResponse) {
	if resp.GetSuccess() {
		UIPrintln("sent.")
	} else {
		UIPrintResponseError(resp, resp.GetResultCode())
	}
}

func (c *Client) handleTellAllNotification(note *message.TellAllNotification) {
	if note.GetSuccess() {
		UIPrintf("tell_all %s> %s", note.Sender, note.Value)
	} else {
		UIPrintResponseError(note, note.GetResultCode())
	}
}
