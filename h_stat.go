package main

import (
	"fmt"
	message "github.com/trasa/watchmud-message"
	"strings"
)

func (c *Client) handleStatResponse(r *message.StatResponse) {
	if !r.GetSuccess() {
		UIPrintResponseError(r, r.GetResultCode())
		return
	}
	var str strings.Builder
	str.WriteString("Status:\n")
	str.WriteString(fmt.Sprintf("Name:\t%s\n", r.PlayerName))
	str.WriteString(fmt.Sprintf("Race:\t%s\t\tClass:\t%s\n", r.GetRace(), r.GetClass()))
	str.WriteString(fmt.Sprintf("Health:\t%d of %d\n", r.CurrentHealth, r.MaxHealth))
	str.WriteString(fmt.Sprintf("Location:\t%s - %s\n", r.ZoneId, r.RoomId))
	str.WriteString(fmt.Sprintf("STR: %d\tDEX: %d\tCON: %d\n", r.Strength, r.Dexterity, r.Constitution))
	str.WriteString(fmt.Sprintf("WIS: %d\tINT: %d\tCHA: %d\n", r.Wisdom, r.Intelligence, r.Charisma))
	str.WriteString("\n")
	UIPrintf(str.String())
}
