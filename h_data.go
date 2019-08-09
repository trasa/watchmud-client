package main

import (
	"encoding/json"
	"errors"
	"github.com/trasa/watchmud-message"
)

func (c *Client) sendDataRequest() error {
	req := message.DataRequest{}
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
	for i, dataType := range resp.DataType {
		if dataType == "races" {
			var races []RaceData
			if err := json.Unmarshal(resp.Data[i], &races); err != nil {
				UIPrintln("failed to unmarshal data: ", err)
				return err
			}
			database.SetRaces(races)
		}
	}
	UIPrintln("Initial Data retrieved successfully.")
	c.clientState.inputHandler = initialInputHandler
	return nil
}
