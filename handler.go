package main

import (
	"github.com/trasa/watchmud-message"
	"log"
)

func (c *Client) handleIncomingMessage(msg *message.GameMessage) error {
	switch msg.Inner.(type) {
	case *message.GameMessage_CreatePlayerResponse:
		if err := c.handleCreatePlayerResponse(msg.GetCreatePlayerResponse()); err != nil {
			log.Printf("Error handling create player response: %v", err)
		}

	case *message.GameMessage_DataResponse:
		if err := c.handleDataResponse(msg.GetDataResponse()); err != nil {
			log.Printf("Error handling data response: %v", err)
		}

	case *message.GameMessage_DeathNotification:
		c.handleDeathNotification(msg.GetDeathNotification())

	case *message.GameMessage_DropNotification:
		c.handleDropNotification(msg.GetDropNotification())

	case *message.GameMessage_DropResponse:
		c.handleDropResponse(msg.GetDropResponse())

	case *message.GameMessage_EnterRoomNotification:
		c.handleEnterRoomNotification(msg.GetEnterRoomNotification())

	case *message.GameMessage_EquipResponse:
		c.handleEquipResponse(msg.GetEquipResponse())

	case *message.GameMessage_ErrorResponse:
		c.handleErrorResponse(msg.GetErrorResponse())

	case *message.GameMessage_ExitsResponse:
		c.handleExitsResponse(msg.GetExitsResponse())

	case *message.GameMessage_GetNotification:
		c.handleGetNotification(msg.GetGetNotification())

	case *message.GameMessage_GetResponse:
		c.handleGetResponse(msg.GetGetResponse())

	case *message.GameMessage_InventoryResponse:
		c.handleInventoryResponse(msg.GetInventoryResponse())

	case *message.GameMessage_KillResponse:
		c.handleKillResponse(msg.GetKillResponse())

	case *message.GameMessage_LeaveRoomNotification:
		c.handleLeaveRoomNotification(msg.GetLeaveRoomNotification())

	case *message.GameMessage_LoginResponse:
		return c.handleLoginResponse(msg.GetLoginResponse())

	case *message.GameMessage_LogoutNotification:
		c.handleLogoutNotification(msg.GetLogoutNotification())

	case *message.GameMessage_LookResponse:
		c.handleLookResponse(msg.GetLookResponse())

	case *message.GameMessage_MoveResponse:
		c.handleMoveResponse(msg.GetMoveResponse())

	case *message.GameMessage_Ping:
		c.handlePing(msg.GetPing())

	case *message.GameMessage_Pong:
		c.handlePong(msg.GetPong())

	case *message.GameMessage_RecallResponse:
		c.handleRecallResponse(msg.GetRecallResponse())

	case *message.GameMessage_RoomStatusResponse:
		c.handleRoomStatusResponse(msg.GetRoomStatusResponse())

	case *message.GameMessage_SayResponse:
		c.handleSayResponse(msg.GetSayResponse())

	case *message.GameMessage_SayNotification:
		c.handleSayNotification(msg.GetSayNotification())

	case *message.GameMessage_ShowEquipmentResponse:
		c.handleShowEquipmentResponse(msg.GetShowEquipmentResponse())

	case *message.GameMessage_TellNotification:
		c.handleTellNotification(msg.GetTellNotification())

	case *message.GameMessage_TellResponse:
		c.handleTellResponse(msg.GetTellResponse())

	case *message.GameMessage_TellAllResponse:
		c.handleTellAllResponse(msg.GetTellAllResponse())

	case *message.GameMessage_TellAllNotification:
		c.handleTellAllNotification(msg.GetTellAllNotification())

	case *message.GameMessage_ViolenceNotification:
		c.handleViolenceNotification(msg.GetViolenceNotification())

	case *message.GameMessage_WearResponse:
		c.handleWearResponse(msg.GetWearResponse())

	case *message.GameMessage_WhoResponse:
		c.handleWhoResponse(msg.GetWhoResponse())

	default:
		log.Printf("client.handleIncomingResponse: unknown response type: %s", msg.Inner)
	}
	return nil
}
