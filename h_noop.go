package main

import (
	"fmt"
	"github.com/trasa/watchmud/message"
)

func (c *Client) handleNoOpResponse(r *message.NoOpResponse) {
	if //noinspection GoBoolExpressions
	VERBOSE_LOGGING {
		fmt.Printf("No Op: success=%t, result=%s\n", r.IsSuccessful(), r.GetResultCode())
	}
}
