package main

import (
	"github.com/trasa/watchmud-message"
	"github.com/trasa/watchmud-message/slot"
)

func (c *Client) handleShowEquipmentResponse(r *message.ShowEquipmentResponse) {
	if !r.GetSuccess() {
		UIPrintError(r, r.GetResultCode())
		return
	}
	UIPrintln("You are using:")
	if len(r.EquipmentInfo) == 0 {
		UIPrintln("Nothing.")
	} else {
		for _, item := range r.EquipmentInfo {
			UIPrintf("%s\t%s\t%s", slot.Location(item.SlotLocation).String(), item.Id, item.ShortDescription)
		}
		UIPrintln()
	}
}
