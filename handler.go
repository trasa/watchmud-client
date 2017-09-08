package main

import (
	"github.com/trasa/watchmud/message"
	"log"
)

func (c *Client) handleIncomingResponse(resp message.Response) {
	switch resp.(type) {
	case *message.EnterRoomNotification:
		c.handleEnterRoomNotification(resp.(*message.EnterRoomNotification))

	case *message.ErrorResponse:
		c.handleErrorResponse(resp.(*message.ErrorResponse))

	case *message.ExitsResponse:
		c.handleExitsResponse(resp.(*message.ExitsResponse))

	case *message.LeaveRoomNotification:
		c.handleLeaveRoomNotification(resp.(*message.LeaveRoomNotification))

	case *message.LoginResponse:
		c.handleLoginResponse(resp.(*message.LoginResponse))

	case *message.LookResponse:
		c.handleLookResponse(resp.(*message.LookResponse))

	case *message.MoveResponse:
		c.handleMoveResponse(resp.(*message.MoveResponse))

	case *message.SayResponse:
		c.handleSayResponse(resp.(*message.SayResponse))

	case *message.SayNotification:
		c.handleSayNotification(resp.(*message.SayNotification))

	default:
		log.Println("unknown response type", resp)
	}
}
