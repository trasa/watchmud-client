package main

type Config struct {
	playerName string
	serverHost string
	serverPort int
}

// global var holding the Config that is in use
var ActiveConfig *Config
