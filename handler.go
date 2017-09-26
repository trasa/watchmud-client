package main

import (
	"github.com/trasa/watchmud/message"
	"log"
)

func (c *Client) handleIncomingResponse(resp message.Response) {
	switch resp.(type) {
	case *message.DropNotification:
		c.handleDropNotification(resp.(*message.DropNotification))

	case *message.DropResponse:
		c.handleDropResponse(resp.(*message.DropResponse))

	case *message.EnterRoomNotification:
		c.handleEnterRoomNotification(resp.(*message.EnterRoomNotification))

	case *message.ErrorResponse:
		c.handleErrorResponse(resp.(*message.ErrorResponse))

	case *message.ExitsResponse:
		c.handleExitsResponse(resp.(*message.ExitsResponse))

	case *message.GetNotification:
		c.handleGetNotification(resp.(*message.GetNotification))

	case *message.GetResponse:
		c.handleGetResponse(resp.(*message.GetResponse))

	case *message.InventoryResponse:
		c.handleInventoryResponse(resp.(*message.InventoryResponse))

	case *message.LeaveRoomNotification:
		c.handleLeaveRoomNotification(resp.(*message.LeaveRoomNotification))

	case *message.LoginResponse:
		c.handleLoginResponse(resp.(*message.LoginResponse))

	case *message.LookResponse:
		c.handleLookResponse(resp.(*message.LookResponse))

	case *message.MoveResponse:
		c.handleMoveResponse(resp.(*message.MoveResponse))

	case *message.NoOpResponse:
		c.handleNoOpResponse(resp.(*message.NoOpResponse))

	case *message.SayResponse:
		c.handleSayResponse(resp.(*message.SayResponse))

	case *message.SayNotification:
		c.handleSayNotification(resp.(*message.SayNotification))

	case *message.TellNotification:
		c.handleTellNotification(resp.(*message.TellNotification))

	case *message.TellResponse:
		c.handleTellResponse(resp.(*message.TellResponse))

	case *message.TellAllResponse:
		c.handleTellAllResponse(resp.(*message.TellAllResponse))

	case *message.TellAllNotification:
		c.handleTellAllNotification(resp.(*message.TellAllNotification))

	case *message.WhoResponse:
		c.handleWhoResponse(resp.(*message.WhoResponse))

	default:
		log.Printf("client.handleIncomingResponse: unknown response type: %s", resp)
	}
}
