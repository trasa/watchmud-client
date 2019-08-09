package main

import (
	"encoding/json"
	"errors"
	"github.com/trasa/watchmud-message"
)

func (c *Client) sendDataRequest() error {
	req := message.DataRequest{
		DataType: "races",
	}
	msg, err := message.NewGameMessage(req)
	if err != nil {
		return err
	}
	c.SendMessage(msg)
	return nil
}

func (c *Client) handleDataResponse(resp *message.DataResponse) error {
	if !resp.GetSuccess() {
		UIPrintln("DataRequest failed: ", resp.GetResultCode())
		return errors.New(resp.GetResultCode())
	}
	UIPrintln("Initial Data retrieved successfully.")
	if resp.DataType == "races" {
		var races []RaceData
		if err := json.Unmarshal(resp.Data, &races); err != nil {
			UIPrintln("failed to unmarshal data: ", err)
			return err
		}
		database.SetRaces(races)
	}
	c.clientState.inputHandler = initialInputHandler
	return nil
}
