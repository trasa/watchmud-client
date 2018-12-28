package main

func createPlayerNameInputHandler(c *Client, tokens []string) {
	UIPrintln("Create a player named", tokens[0])
	c.clientState.playerName = tokens[0]

	if err := c.sendCreatePlayerRequest(); err != nil {
		UIPrintError(err)
		UIPrintln("Let's try again. Whats your name?")
		// same handler
	} else {
		// otherwise we're waiting for the CreatePlayerResponse
		c.clientState.inputHandler = voidInputHandler
	}
}
