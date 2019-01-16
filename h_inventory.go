package main

import (
	"fmt"
	"github.com/trasa/watchmud-message"
	"strings"
)

func (c *Client) handleInventoryResponse(r *message.InventoryResponse) {
	if !r.GetSuccess() {
		UIPrintResponseError(r, r.GetResultCode())
		return
	}
	var str strings.Builder
	str.WriteString("You are carrying:\n")
	if len(r.InventoryItems) == 0 {
		str.WriteString(" Nothing.\n")
	} else {
		for _, item := range r.InventoryItems {
			str.WriteString(fmt.Sprintf("%s\t%s\n", item.Id, item.ShortDescription))
		}
		str.WriteString("\n")
	}
	UIPrintf(str.String())
}
