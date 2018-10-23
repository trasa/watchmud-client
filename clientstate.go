package main

/*

initial_state:
	choices: login, create_player


login:
	prompt_for_player_name
	prompt_for_player_password
	send_login_request
	receive_login_response
	if response==success: goto in_game
	else goto initial_state

create_player:
	prompt_for_player_name
	prompt_for_player_password
	[TODO prompt for other details]
	send create_player_request
	get create_player_response
	if response==success: goto in_game
	else goto initial_state
*/

type GameState int32

//go:generate stringer -type=GameState
const (
	Initial GameState = iota
	Login
	CreatePlayer
	InGame
)

type ClientState struct {
	currentState GameState
	playerName   string
}

func NewClientState() *ClientState {
	return &ClientState{
		currentState: Initial,
	}
}

/*
Moves from Login -> Initial
*/
func (cs *ClientState) handleLoginFailed() {
	if cs.currentState != Login {
		UIPrintf("ClientState received handleLoginFailed but is in %s state\n", cs.currentState)
	}
	cs.currentState = Initial
	UIPrintln("(DEBUG) currentState now Initial")
}

/*
Moves from Login -> InGame
*/
func (cs *ClientState) handleLoginSucceeded(playerName string) {
	if cs.currentState != Login {
		UIPrintf("ClientState received handleLoginSucceeded but is in %s state\n", cs.currentState)
	}
	cs.playerName = playerName
	cs.currentState = InGame
	UIPrintln("(DEBUG) currentState now InGame")
}
