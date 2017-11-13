package main

import (
	"fmt"
	"github.com/trasa/watchmud/direction"
	"github.com/trasa/watchmud/message"
	"reflect"
)

func (c *Client) printPrompt() {
	// TODO need to figure out when the right time to print the prompt is ...
	fmt.Print("> ")
}

// for a response with IsSuccess == false,
// print a generic error message.
func (c *Client) printError(response interface{}, resultCode string) {
	fmt.Println("Error:", reflect.TypeOf(response).String(), resultCode)
}

// print this room description to the player
func (c *Client) printRoom(room *message.RoomDescription) {
	fmt.Println(room.Name)
	fmt.Println()
	fmt.Println(room.Description)
	// obvious exits
	if exits, err := direction.ExitsToFormattedString(room.Exits); err == nil {
		fmt.Println("Obvious Exits:", exits)
	} else {
		fmt.Println("Error Getting exits:", err)
	}
	if len(room.Players) > 0 || len(room.Objects) > 0 {
		fmt.Println()
	}
	// objects
	for _, o := range room.Objects {
		fmt.Println(o)
	}
	// other players
	for _, m := range room.Mobs {
		fmt.Println(m)
	}
	for _, p := range room.Players {
		fmt.Printf("%s stands here.\n", p)
	}
}
