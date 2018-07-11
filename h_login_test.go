package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/trasa/watchmud-message"
	"testing"
)

func TestLoginResponse_success(t *testing.T) {
	c := NewTestClient()
	resp := &message.LoginResponse{
		Success:    true,
		PlayerName: "testguy",
	}
	c.handleLoginResponse(resp)
	assert.Equal(t, "testguy", c.playerName)
}

func TestLoginResponse_loginFailed(t *testing.T) {
	c := NewTestClient()
	resp := &message.LoginResponse{
		Success:    false,
		ResultCode: "BAD_DATA",
	}
	c.handleLoginResponse(resp)
	assert.Equal(t, "", c.playerName)
}
