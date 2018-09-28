package main

import (
	"github.com/trasa/watchmud-message"
)

func (c *Client) handleInventoryResponse(r *message.InventoryResponse) {
	if !r.GetSuccess() {
		UIPrintError(r, r.GetResultCode())
		return
	}
	UIPrintln("You are carrying:")
	if len(r.InventoryItems) == 0 {
		UIPrintln(" Nothing.")
	} else {
		for _, item := range r.InventoryItems {
			UIPrintf("%s\t%s\n", item.Id, item.ShortDescription)
		}
		UIPrintln()
	}
}
