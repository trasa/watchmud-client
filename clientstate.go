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

type InputHandler func(c *Client, tokens []string)

type ClientState struct {
	playerName   string
	inputHandler InputHandler
	race         int32
	class        int32
}

func NewClientState() *ClientState {
	return &ClientState{
		inputHandler: initialInputHandler,
	}
}

// The initial InputHandler for ClientState. Move from here to login or create.
func initialInputHandler(c *Client, tokens []string) {
	switch tokens[0] {
	case "login":
		// next step is to get player name - do we have one already?
		if len(tokens) < 2 {
			UIPrintln("Who are you?")
			// set input handler for 'get player name'
			c.clientState.inputHandler = loginNameInputHandler
		} else {
			c.clientState.playerName = tokens[1]
			if err := c.sendLoginRequest(); err != nil {
				UIPrintError(err)
				// keep same input handler
			}
		}
		break
	case "create":
		UIPrintln("Ok, let's create a new character. What is your name?")
		c.clientState.inputHandler = createPlayerNameInputHandler
		break

	case "help":
		// TODO context sensitive help
		printHelp(tokens)
		// keep same handler

	default:
		UIPrintln("Unknown command, try 'login' or 'create' or 'help'")
		// keep same handler
	}
}

// Default game handler, do the normal stuff with the
// input as if you were in the game.
func gameInputHandler(c *Client, tokens []string) {
	c.sendTokens(tokens)
}

// Game handler when we're not allowing input from the user
// probably because we're waiting for a response message from
// the server
func voidInputHandler(c *Client, tokens []string) {
	UIPrintln("Please wait ...")
}
