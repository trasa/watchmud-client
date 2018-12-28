package main

// login states:
// now we have a name, lets send a login request
func loginNameInputHandler(c *Client, tokens []string) {
	UIPrintln("Hello '", tokens[0], "'")
	c.clientState.playerName = tokens[0]
	if err := c.sendLoginRequest(); err != nil {
		UIPrintError(err)
		UIPrintln("No really, who are you?")
		// same handler
	} else {
		// otherwise waiting for response
		c.clientState.inputHandler = voidInputHandler
	}
}
