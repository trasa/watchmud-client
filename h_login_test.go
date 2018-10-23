package main

import (
	"github.com/stretchr/testify/suite"
	"github.com/trasa/watchmud-message"
	"testing"
)

type LoginResponseSuite struct {
	suite.Suite
	*Client
}

func TestLoginResponseSuite(t *testing.T) {
	suite.Run(t, new(LoginResponseSuite))
}

func (suite *LoginResponseSuite) SetUpTest() {
	suite.Client = NewTestClient()
	activeClientUI = NewClientUI(suite.Client)
	activeClientUI.isTesting = true
}

func (suite *LoginResponseSuite) success() {
	resp := &message.LoginResponse{
		Success:    true,
		PlayerName: "testguy",
	}
	suite.Client.handleLoginResponse(resp)
	suite.Assert().Equal("testguy", suite.Client.clientState.playerName)
}

func (suite *LoginResponseSuite) loginFailed() {
	resp := &message.LoginResponse{
		Success:    false,
		ResultCode: "BAD_DATA",
	}
	suite.Client.handleLoginResponse(resp)
	suite.Assert().Equal("", suite.Client.clientState.playerName)
}
