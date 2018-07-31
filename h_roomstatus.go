package main

import (
	"fmt"
	"github.com/trasa/watchmud-message"
)

func (c *Client) handleRoomStatusResponse(resp *message.RoomStatusResponse) {
	if !resp.Success {
		c.printError(resp, resp.ResultCode)
		return
	}

	fmt.Println()
	fmt.Printf("-----------------------------------------------------\n")
	fmt.Printf("ROOM STATUS\n")
	fmt.Printf("-----------------------------------------------------\n")

	fmt.Printf("Room %s (%s)\nZone %s (%s)\n", resp.Name, resp.Id, resp.ZoneName, resp.ZoneId)
	fmt.Printf("'%s'\n", resp.Description)
	fmt.Printf("\n")

	if len(resp.GetDirections()) > 0 {
		fmt.Printf("Exit to Room\n")
		fmt.Printf("-----------------------------------------------------\n")
		for _, d := range resp.GetDirections() {
			fmt.Printf("%-7s %-20s %-10s  ", d.Dir, d.RoomId, d.ZoneId)
			c.printStringList(d.Flags)
			fmt.Printf("\n")
		}
		fmt.Println()
	}
	if len(resp.PlayerInfo) > 0 {
		fmt.Printf("Players\n")
		fmt.Printf("-----------------------------------------------------\n")
		for _, p := range resp.PlayerInfo {
			fmt.Printf("%-16s(%d/%d)\n", p.Name, p.CurrentHealth, p.MaxHealth)
			fmt.Printf("-------------------------\n")
		}
		fmt.Println()
	}
	if len(resp.MobInfo) > 0 {
		fmt.Printf("Mobs\n")
		fmt.Printf("-----------------------------------------------------\n")
		for _, m := range resp.MobInfo {
			fmt.Printf("%s (%d/%d) Aliases: ", m.Name, m.CurrentHealth, m.MaxHealth)
			c.printStringList(m.Aliases)
			fmt.Printf("\n")
			fmt.Printf("ID: %s\n", m.Id)
			fmt.Printf("Defn: %s\n", m.DefinitionId)
			fmt.Printf("Zone: %s\n", m.ZoneId)
			fmt.Printf("'%s'\n", m.DescriptionInRoom)
			fmt.Printf("Descr:\t%s\n", m.ShortDescription)
			fmt.Printf("Flags: ")
			c.printStringList(m.Flags)
			fmt.Printf("\n")
			fmt.Printf("-------------------------\n")
		}
		fmt.Println()
	}
	if len(resp.InventoryInfo) > 0 {
		fmt.Printf("Items\n")
		fmt.Printf("-----------------------------------------------------\n")
		for _, item := range resp.InventoryInfo {
			fmt.Printf("%-16s Aliases: ", item.Name)
			c.printStringList(item.Aliases)
			fmt.Printf("\n")
			fmt.Printf("ID: %s\n", item.Id)
			fmt.Printf("Defn: %s\n", item.DefinitionId)
			fmt.Printf("Zone: %s\n", item.ZoneId)
			fmt.Printf("Categories: ")
			c.printStringList(item.Categories)
			fmt.Printf("\n")
			fmt.Printf("Descr: %s\n", item.ShortDescription)
			fmt.Printf("On Ground: %s\n", item.DescriptionOnGround)
			fmt.Printf("Behaviors: ")
			c.printStringList(item.Behaviors)
			fmt.Printf("\n")
			fmt.Printf("-------------------------\n")
		}
		fmt.Println()
	}
	fmt.Printf("\n")
}
