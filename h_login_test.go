package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/trasa/watchmud/message"
	"testing"
)

func TestLoginResponse_success(t *testing.T) {
	c := NewTestClient()
	resp := &message.LoginResponse{
		Response: message.NewSuccessfulResponse("login_response"),
		Player: message.PlayerData{
			Name: "testguy",
		},
	}
	c.handleLoginResponse(resp)
	assert.Equal(t, "testguy", c.playerData.Name)
}

func TestLoginResponse_loginFailed(t *testing.T) {
	c := NewTestClient()
	resp := &message.LoginResponse{
		Response: message.NewUnsuccessfulResponse("login_response", "BAD_DATA"),
	}
	c.handleLoginResponse(resp)
	assert.Equal(t, "", c.playerData.Name)
}
