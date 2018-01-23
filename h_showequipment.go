package main

import (
	"fmt"
	"github.com/trasa/watchmud/message"
	"github.com/trasa/watchmud/slot"
)

func (c *Client) handleShowEquipmentResponse(r *message.ShowEquipmentResponse) {
	if !r.GetSuccess() {
		c.printError(r, r.GetResultCode())
		return
	}
	fmt.Println("You are using:")
	if len(r.EquipmentInfo) == 0 {
		fmt.Println("Nothing.")
	} else {
		for _, item := range r.EquipmentInfo {
			fmt.Printf("%s\t%s\t%s", slot.Location(item.SlotLocation).String(), item.Id, item.ShortDescription)
		}
		fmt.Println()
	}
}
