package main

type Config struct {
	serverHost string
	serverPort int
}

// global var holding the Config that is in use
var ActiveConfig *Config
