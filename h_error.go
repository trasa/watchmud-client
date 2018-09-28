package main

import "github.com/trasa/watchmud-message"

func (c *Client) handleErrorResponse(resp *message.ErrorResponse) {
	UIPrintError(resp, resp.GetResultCode())
}
