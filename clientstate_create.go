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
}

func displayRaceChoices() {
	UIPrintln("Select a Race:\n" +
		"0 - Human\n" +
		"1 - Dwarf\n" +
		"2 - Elf\n" +
		"3 - Halfling\n" +
		"4 - Dragonborn\n" +
		"5 - Half-Elf\n" +
		"6 - Half-Orc\n" +
		"7 - Gnome\n" +
		"8 - Tiefling\n" +
		"\n" +
		"(or help for more information)\n" +
		"ex. HELP DWARF")
}

// select a race
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
	UIPrintln("Select a class:\n" +
		"0 - Fighter\n" +
		"1 - Cleric\n" +
		"2 - Rogue\n" +
		"3 - Barbarian\n" +
		"4 - Bard\n" +
		"5 - Druid\n" +
		"6 - Monk\n" +
		"7 - Paladin\n" +
		"8 - Ranger\n" +
		"9 - Sorcerer\n" +
		"10 - Warlock\n" +
		"11 - Wizard\n" +
		"\n" +
		"(or help for more information)" +
		"ex. HELP FIGHTER")
}

// select a class
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
