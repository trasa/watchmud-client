package main

import (
	"encoding/json"
	"errors"
	"github.com/trasa/watchmud-message"
	"strings"
)

type RaceData struct {
	RaceId        int32
	RaceGroupName string
	RaceName      string
}

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
	var races []RaceData
	if err := json.Unmarshal(resp.Data, &races); err != nil {
		UIPrintln("failed to unmarshal data: ", err)
		return err
	}
	// TODO just for debugging
	var str strings.Builder
	for _, r := range races {
		str.WriteString(r.RaceGroupName + " " + r.RaceName)
	}
	UIPrintln(str.String())
	c.clientState.inputHandler = initialInputHandler
	return nil
}
