package main

// prints help to stdout
func printHelp(tokens []string) {
	if len(tokens) <= 1 {
		// print commands
		UIPrintln("TODO help for all commands goes here.")
	} else {
		// figure out which command we're asking for help for
		UIPrintf("TODO help for '%s' goes here.\n", tokens[1])
	}
}
