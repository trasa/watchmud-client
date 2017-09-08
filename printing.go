package main

import (
	"fmt"
	"github.com/trasa/watchmud/direction"
	"github.com/trasa/watchmud/message"
)

func (c *Client) printPrompt() {
	// TODO need to figure out when the right time to print the prompt is ...
	fmt.Print("> ")
}

// for a response with IsSuccess == false,
// print a generic error message.
func (c *Client) printError(resp message.Response) {
	fmt.Println("Error:", resp.GetMessageType(), resp.GetResultCode())
}

// print this room description to the player
func (c *Client) printRoom(room message.RoomDescription) {
	fmt.Println(room.Name)
	fmt.Println()
	fmt.Println(room.Description)
	// other players
	if len(room.Players) > 0 {
		fmt.Println("Other Players:")
		for _, p := range room.Players {
			fmt.Println(p)
		}
	}
	// obvious exits
	if exits, err := direction.ExitsToFormattedString(room.Exits); err == nil {
		fmt.Println("Obvious Exits:", exits)
	} else {
		fmt.Println("Error Getting exits:", err)
	}
}
