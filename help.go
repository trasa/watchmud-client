package main

import "fmt"

// prints help to stdout
func printHelp(tokens []string) {
	if len(tokens) <= 1 {
		// print commands
		fmt.Println("TODO help for all commands goes here.")
	} else {
		// figure out which command we're asking for help for
		fmt.Printf("TODO help for '%s' goes here.\n", tokens[1])
	}
}
