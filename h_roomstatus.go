package main

import (
	"fmt"
	"github.com/trasa/watchmud-message"
	"strings"
)

func (c *Client) handleRoomStatusResponse(resp *message.RoomStatusResponse) {
	if !resp.Success {
		UIPrintResponseError(resp, resp.ResultCode)
		return
	}

	var str strings.Builder
	str.WriteString("\n-----------------------------------------------------\n")
	str.WriteString("ROOM STATUS\n")
	str.WriteString("-----------------------------------------------------\n")

	str.WriteString(fmt.Sprintf("Room %s (%s)\nZone %s (%s)\n", resp.Name, resp.Id, resp.ZoneName, resp.ZoneId))
	str.WriteString(fmt.Sprintf("'%s'\n", resp.Description))
	str.WriteString("\n")

	if len(resp.GetDirections()) > 0 {
		str.WriteString("Exit to Room\n")
		str.WriteString("-----------------------------------------------------\n")
		for _, d := range resp.GetDirections() {
			str.WriteString(fmt.Sprintf("%-7s %-20s %-10s  ", d.Dir, d.RoomId, d.ZoneId))
			str.WriteString(ListToStringBuilder(d.Flags))
			str.WriteString("\n")
		}
		str.WriteString("\n")
	}
	if len(resp.PlayerInfo) > 0 {
		str.WriteString("Players\n")
		str.WriteString("-----------------------------------------------------\n")
		for _, p := range resp.PlayerInfo {
			str.WriteString(fmt.Sprintf("%-16s(%d/%d)\n", p.Name, p.CurrentHealth, p.MaxHealth))
			str.WriteString("-----------------------------------------------------\n")
		}
		str.WriteString("\n")
	}
	if len(resp.MobInfo) > 0 {
		str.WriteString("Mobs\n")
		str.WriteString("-----------------------------------------------------\n")
		for _, m := range resp.MobInfo {
			str.WriteString(fmt.Sprintf("%s (%d/%d) Aliases: ", m.Name, m.CurrentHealth, m.MaxHealth))
			str.WriteString(ListToStringBuilder(m.Aliases))
			str.WriteString("\n")
			str.WriteString(fmt.Sprintf("ID: %s\n", m.Id))
			str.WriteString(fmt.Sprintf("Defn: %s\n", m.DefinitionId))
			str.WriteString(fmt.Sprintf("Zone: %s\n", m.ZoneId))
			str.WriteString(fmt.Sprintf("'%s'\n", m.DescriptionInRoom))
			str.WriteString(fmt.Sprintf("Descr:\t%s\n", m.ShortDescription))
			str.WriteString("Flags: ")
			str.WriteString(ListToStringBuilder(m.Flags))
			str.WriteString("\n")
			str.WriteString("-----------------------------------------------------\n")
		}
		str.WriteString("\n")
	}
	if len(resp.InventoryInfo) > 0 {
		str.WriteString("Items\n")
		str.WriteString("-----------------------------------------------------\n")
		for _, item := range resp.InventoryInfo {
			str.WriteString(fmt.Sprintf("%-16s Aliases: ", item.Name))
			str.WriteString(ListToStringBuilder(item.Aliases))
			str.WriteString("\n")
			str.WriteString(fmt.Sprintf("ID: %s\n", item.Id))
			str.WriteString(fmt.Sprintf("Defn: %s\n", item.DefinitionId))
			str.WriteString(fmt.Sprintf("Zone: %s\n", item.ZoneId))
			str.WriteString("Categories: ")
			str.WriteString(ListToStringBuilder(item.Categories))
			str.WriteString("\n")
			str.WriteString(fmt.Sprintf("Descr: %s\n", item.ShortDescription))
			str.WriteString(fmt.Sprintf("On Ground: %s\n", item.DescriptionOnGround))
			str.WriteString("Behaviors: ")
			str.WriteString(ListToStringBuilder(item.Behaviors))
			str.WriteString("\n")
			str.WriteString("-----------------------------------------------------\n")
		}
		str.WriteString("\n")
	}
	str.WriteString("\n")

	UIPrint(str)
}
