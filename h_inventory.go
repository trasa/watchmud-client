package main

import (
	"fmt"
	"github.com/trasa/watchmud-message"
)

func (c *Client) handleInventoryResponse(r *message.InventoryResponse) {
	if !r.GetSuccess() {
		c.printError(r, r.GetResultCode())
		return
	}
	fmt.Println("You are carrying:")
	if len(r.InventoryItems) == 0 {
		fmt.Println(" Nothing.")
	} else {
		for _, item := range r.InventoryItems {
			fmt.Printf("%s\t%s\n", item.Id, item.ShortDescription)
		}
		fmt.Println()
	}
}
