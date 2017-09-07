package main

import (
	"github.com/trasa/watchmud/message"
	"log"
)

func (c *Client) handleIncomingResponse(resp message.Response) {
	switch resp.(type) {
	case *message.EnterRoomNotification:
		c.handleEnterRoomNotification(resp.(*message.EnterRoomNotification))

	case *message.ExitsResponse:
		c.handleExitsResponse(resp.(*message.ExitsResponse))

	case *message.LoginResponse:
		c.handleLoginResponse(resp.(*message.LoginResponse))

	case *message.LookResponse:
		c.handleLookResponse(resp.(*message.LookResponse))

	case *message.MoveResponse:
		c.handleMoveResponse(resp.(*message.MoveResponse))

	default:
		log.Println("unknown response type", resp)
	}
}
