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
			} else {
				// set next step input handler
			}
		}
		break
	case "create":
		// TODO create player not implemented yet
		UIPrintln("Create Player not implemented yet, try 'login' or 'help'")
		// keep same handler
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

func loginNameInputHandler(c *Client, tokens []string) {
	UIPrintln("Hello '", tokens[0], "'")
	c.clientState.playerName = tokens[0]
	if err := c.sendLoginRequest(); err != nil {
		UIPrintError(err)
		UIPrintln("No really, who are you?")
		// same handler
	} else {
		// ok now what?
	}
}

func gameInputHandler(c *Client, tokens []string) {
	c.sendTokens(tokens)
}
