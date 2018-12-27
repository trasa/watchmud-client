package main

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TargetingSuite struct {
	suite.Suite
	testClient *Client
}

func TestTargetingSuite(t *testing.T) {
	suite.Run(t, new(TargetingSuite))
}

func (suite *TargetingSuite) SetupTest() {
	suite.testClient.clientState.playerName = "testdood"
}

func (suite *TargetingSuite) TargetMatchesPlayerName() {
	suite.Assert().True(suite.testClient.IsTargetYou("testdood"))
}

func (suite *TargetingSuite) TargetDoesNotMatch() {
	suite.Assert().True(suite.testClient.IsTargetYou("otherguy"))
}
