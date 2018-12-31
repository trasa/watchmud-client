package main

import (
	"strconv"
	"strings"
)

// "OK, what is your name?"
// > bob
func createPlayerNameInputHandler(c *Client, tokens []string) {
	UIPrintln("Create a player named", tokens[0])
	c.clientState.playerName = tokens[0]

	displayRaceChoices()
	c.clientState.inputHandler = createPlayerRaceInputHandler

	/*
		if err := c.sendCreatePlayerRequest(); err != nil {
			UIPrintError(err)
			UIPrintln("Let's try again. Whats your name?")
			// same handler
		} else {
			// otherwise we're waiting for the CreatePlayerResponse
			c.clientState.inputHandler = voidInputHandler
		}
	*/
}

func displayRaceChoices() {
	UIPrintln("Select a Race:")
	UIPrintln("0 - Human")
	UIPrintln("1 - Dwarf")
	UIPrintln("2 - Elf")
	UIPrintln("3 - Halfling")
	UIPrintln("4 - Dragonborn")
	UIPrintln("5 - Half-Elf")
	UIPrintln("6 - Half-Orc")
	UIPrintln("7 - Gnome")
	UIPrintln("8 - Tiefling")
	UIPrintln()
	UIPrintln("(or help for more information)")
	UIPrintln("ex. HELP DWARF")
}

func createPlayerRaceInputHandler(c *Client, tokens []string) {
	if strings.EqualFold("help", tokens[0]) {
		UIPrintln("Help not implemented yet, sorry...") // TODO implement help..
		displayRaceChoices()                            // same handler
	} else {
		choice, err := strconv.Atoi(tokens[0])
		if err != nil || choice < 0 || choice > 8 {
			UIPrintln("Please select a race from the list.")
			displayRaceChoices()
		} else {
			// set the clientState choice here and move on to the next step
			c.clientState.race = int32(choice)
			displayClassChoices()
			c.clientState.inputHandler = createPlayerClassInputHandler
		}
	}
}

func displayClassChoices() {
	UIPrintln("Select a class:")
	UIPrintln("0 - Fighter")
	UIPrintln("1 - Cleric")
	UIPrintln("2 - Rogue")
	UIPrintln("3 - Barbarian")
	UIPrintln("4 - Bard")
	UIPrintln("5 - Druid")
	UIPrintln("6 - Monk")
	UIPrintln("7 - Paladin")
	UIPrintln("8 - Ranger")
	UIPrintln("9 - Sorcerer")
	UIPrintln("10 - Warlock")
	UIPrintln("11 - Wizard")
	UIPrintln()
	UIPrintln("(or help for more information)")
	UIPrintln("ex. HELP FIGHTER")

}

func createPlayerClassInputHandler(c *Client, tokens []string) {
	if strings.EqualFold("help", tokens[0]) {
		UIPrintln("Help not implemented yet, sorry...") // TODO implement help...
		displayClassChoices()                           // same handler
	} else {
		choice, err := strconv.Atoi(tokens[0])
		if err != nil || choice < 0 || choice > 11 {
			UIPrintln("Please select a class from the list.")
			displayClassChoices()
		} else {
			c.clientState.class = int32(choice)
			// for now send a message
			c.clientState.inputHandler = voidInputHandler
			c.sendCreatePlayerRequest()
		}
	}
}
