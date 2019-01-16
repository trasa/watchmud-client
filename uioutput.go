package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/trasa/watchmud-message"
	"github.com/trasa/watchmud-message/direction"
	"reflect"
	"strings"
)

func UIPrintError(err error) {
	UIPrintln("Error: ", err)
}

// for a response with IsSuccess == false,
// print a generic error message.
func UIPrintResponseError(response interface{}, resultCode string) {
	UIPrintln("Error:", reflect.TypeOf(response).String(), resultCode)
}

// print this room description to the player
func UIPrintRoom(room *message.RoomDescription) {
	var str strings.Builder

	str.WriteString(room.Name + "\n")
	str.WriteString(room.Description + "\n");
	// obvious exits
	if exits, err := direction.ExitsToFormattedString(room.Exits); err == nil {
		str.WriteString("Obvious Exits: ")
		str.WriteString(exits)
		str.WriteString("\n")

	} else {
		str.WriteString("Error getting exits:")
		str.WriteString(fmt.Sprintf("%v\n", err))
	}
	if len(room.Players) > 0 || len(room.Objects) > 0 {
		str.WriteString("\n")
	}
	// objects
	for _, o := range room.Objects {
		str.WriteString(o + "\n")
	}
	// other players
	for _, m := range room.Mobs {
		str.WriteString(m + "\n")
	}
	for _, p := range room.Players {
		str.WriteString(fmt.Sprintf("%s stands here.\n", p))
	}

	UIPrintf(str.String())
}

// Print an array of strings (flags, aliases, that sort of thing)
func UIPrintStringList(list []string) {
	var str strings.Builder
	flagLen := len(list)
	if flagLen > 0 {
		str.WriteString("[")
		for idx, f := range list {
			str.WriteString(f)
			if idx < flagLen-1 {
				str.WriteString(", ")
			}
		}
		str.WriteString("]")

		UIPrintf(str.String())
	}
}

func UIPrintln(a ...interface{}) {
	if activeClientUI.isTesting {
		fmt.Println(a...)
	} else {
		activeClientUI.g.Update(func(g *gocui.Gui) error {
			_, err := fmt.Fprintln(activeClientUI.outputView(), a...)
			return err
		})
	}
}

func UIPrintf(format string, a ...interface{}) {
	if activeClientUI.isTesting {
		fmt.Printf(format, a...)
	} else {
		activeClientUI.g.Update(func(g *gocui.Gui) error {
			_, err := fmt.Fprintf(activeClientUI.outputView(), format, a...)
			return err
		})
	}
}
