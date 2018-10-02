package main

import (
	"github.com/trasa/watchmud-message"
)

func (c *Client) handleRoomStatusResponse(resp *message.RoomStatusResponse) {
	if !resp.Success {
		UIPrintResponseError(resp, resp.ResultCode)
		return
	}

	UIPrintln()
	UIPrintf("-----------------------------------------------------\n")
	UIPrintf("ROOM STATUS\n")
	UIPrintf("-----------------------------------------------------\n")

	UIPrintf("Room %s (%s)\nZone %s (%s)\n", resp.Name, resp.Id, resp.ZoneName, resp.ZoneId)
	UIPrintf("'%s'\n", resp.Description)
	UIPrintf("\n")

	if len(resp.GetDirections()) > 0 {
		UIPrintf("Exit to Room\n")
		UIPrintf("-----------------------------------------------------\n")
		for _, d := range resp.GetDirections() {
			UIPrintf("%-7s %-20s %-10s  ", d.Dir, d.RoomId, d.ZoneId)
			UIPrintStringList(d.Flags)
			UIPrintf("\n")
		}
		UIPrintln()
	}
	if len(resp.PlayerInfo) > 0 {
		UIPrintf("Players\n")
		UIPrintf("-----------------------------------------------------\n")
		for _, p := range resp.PlayerInfo {
			UIPrintf("%-16s(%d/%d)\n", p.Name, p.CurrentHealth, p.MaxHealth)
			UIPrintf("-----------------------------------------------------\n")
		}
		UIPrintln()
	}
	if len(resp.MobInfo) > 0 {
		UIPrintf("Mobs\n")
		UIPrintf("-----------------------------------------------------\n")
		for _, m := range resp.MobInfo {
			UIPrintf("%s (%d/%d) Aliases: ", m.Name, m.CurrentHealth, m.MaxHealth)
			UIPrintStringList(m.Aliases)
			UIPrintf("\n")
			UIPrintf("ID: %s\n", m.Id)
			UIPrintf("Defn: %s\n", m.DefinitionId)
			UIPrintf("Zone: %s\n", m.ZoneId)
			UIPrintf("'%s'\n", m.DescriptionInRoom)
			UIPrintf("Descr:\t%s\n", m.ShortDescription)
			UIPrintf("Flags: ")
			UIPrintStringList(m.Flags)
			UIPrintf("\n")
			UIPrintf("-----------------------------------------------------\n")
		}
		UIPrintln()
	}
	if len(resp.InventoryInfo) > 0 {
		UIPrintf("Items\n")
		UIPrintf("-----------------------------------------------------\n")
		for _, item := range resp.InventoryInfo {
			UIPrintf("%-16s Aliases: ", item.Name)
			UIPrintStringList(item.Aliases)
			UIPrintf("\n")
			UIPrintf("ID: %s\n", item.Id)
			UIPrintf("Defn: %s\n", item.DefinitionId)
			UIPrintf("Zone: %s\n", item.ZoneId)
			UIPrintf("Categories: ")
			UIPrintStringList(item.Categories)
			UIPrintf("\n")
			UIPrintf("Descr: %s\n", item.ShortDescription)
			UIPrintf("On Ground: %s\n", item.DescriptionOnGround)
			UIPrintf("Behaviors: ")
			UIPrintStringList(item.Behaviors)
			UIPrintf("\n")
			UIPrintf("-----------------------------------------------------\n")
		}
		UIPrintln()
	}
	UIPrintf("\n")
}
