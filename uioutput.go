package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/trasa/watchmud-message"
	"github.com/trasa/watchmud-message/direction"
	"reflect"
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
	UIPrintln(room.Name)
	UIPrintln()
	UIPrintln(room.Description)
	// obvious exits
	if exits, err := direction.ExitsToFormattedString(room.Exits); err == nil {
		UIPrintln("Obvious Exits:", exits)
	} else {
		UIPrintln("Error Getting exits:", err)
	}
	if len(room.Players) > 0 || len(room.Objects) > 0 {
		UIPrintln()
	}
	// objects
	for _, o := range room.Objects {
		UIPrintln(o)
	}
	// other players
	for _, m := range room.Mobs {
		UIPrintln(m)
	}
	for _, p := range room.Players {
		UIPrintf("%s stands here.\n", p)
	}
}

// Print an array of strings (flags, aliases, that sort of thing)
func UIPrintStringList(list []string) {
	flagLen := len(list)
	if flagLen > 0 {
		UIPrintf("[")
		for idx, f := range list {
			UIPrintf("%s", f)
			if idx < flagLen-1 {
				UIPrintf(", ")
			}
		}
		UIPrintf("]")
	}
}

func UIPrintln(a ...interface{}) {
	if activeClientUI.isTesting {
		fmt.Println(a...)
	} else {
		activeClientUI.g.Update(func(g *gocui.Gui) error {
			fmt.Fprintln(activeClientUI.outputView(), a...)
			return nil
		})
	}
}

func UIPrintf(format string, a ...interface{}) {
	if activeClientUI.isTesting {
		fmt.Printf(format, a...)
	} else {
		activeClientUI.g.Update(func(g *gocui.Gui) error {
			fmt.Fprintf(activeClientUI.outputView(), format, a...)
			return nil
		})
	}
}
