package main

import (
	"fmt"
	"github.com/trasa/watchmud-message"
	"github.com/trasa/watchmud-message/slot"
	"strings"
)

func (c *Client) handleShowEquipmentResponse(r *message.ShowEquipmentResponse) {
	if !r.GetSuccess() {
		UIPrintResponseError(r, r.GetResultCode())
		return
	}
	var str strings.Builder
	str.WriteString("You are using:\n")
	if len(r.EquipmentInfo) == 0 {
		str.WriteString("Nothing.\n")
	} else {
		for _, item := range r.EquipmentInfo {
			str.WriteString(fmt.Sprintf("%s\t%s\t%s\n", slot.Location(item.SlotLocation).String(), item.Id, item.ShortDescription))
		}
		str.WriteString("\n")
	}
	UIPrint(str)
}
