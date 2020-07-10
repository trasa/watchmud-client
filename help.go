package main

import "github.com/trasa/watchmud-client/box"

// prints help to stdout
func printHelp(tokens []string) {
	if len(tokens) <= 1 {
		// print commands
		b, success := box.Get("/help.md")
		if !success {
			UIPrintln("Error: Help file 'help.md' not found!")
		} else {
			UIPrintln(string(b))
		}
	} else {
		// figure out which command we're asking for help for
		UIPrintf("TODO help for '%s' goes here.\n", tokens[1])
	}
}
