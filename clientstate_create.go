package main

import (
	"fmt"
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
	var str strings.Builder
	str.WriteString("Select a Race:\n")
	for i := 0; i < len(database.Races); i++ {
		str.WriteString(fmt.Sprintf("%d - %s\n", i, database.Races[int32(i)].RaceName))
	}
	str.WriteString("\n")
	str.WriteString("(or help for more information)\n")
	str.WriteString("ex. HELP DWARF")
	UIPrint(str)
}

// select a race
func createPlayerRaceInputHandler(c *Client, tokens []string) {
	if strings.EqualFold("help", tokens[0]) {
		UIPrintln("Help not implemented yet, sorry...") // TODO implement help..
		displayRaceChoices()                            // same handler
	} else {
		choice, err := strconv.Atoi(tokens[0])
		if err != nil || choice < 0 || choice >= len(database.Races) {
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
