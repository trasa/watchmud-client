package main

func (c *Client) IsTargetYou(target string) bool {
	// TODO get better about comparing targets which are Mob IDs or
	// other IDs vs. player IDs.
	return c.playerName == target
}
