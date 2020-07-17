package main

import (
	"flag"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\t%s [flags]\n", os.Args[0])
	fmt.Fprint(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	// TODO read from yaml configuration or something

	host := flag.String("host", "localhost", "server host name")
	port := flag.Int("port", 10000, "server port")
	doHelp := flag.Bool("help", false, "Show Help")
	doHelpAlias := flag.Bool("h", false, "Show Help")
	logFile := flag.String("logFile", "./watchmud-client.log", "File to write client logs to")
	debug := flag.Bool("debug", false, "Set log level to debug")
	flag.Parse()
	if *doHelp || *doHelpAlias {
		usage()
		os.Exit(2)
		return
	}

	// init logging
	f, err := os.OpenFile(*logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Panic().Msgf("error opening log file: %v", err)
	}
	defer f.Close()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: f})
	log.Info().Msg("Logging initialized.")

	ActiveConfig = &Config{
		serverHost: *host,
		serverPort: *port,
	}

	// connect client
	client, err := Connect(ActiveConfig.serverHost, ActiveConfig.serverPort)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect")
	}
	signal.Notify(client.quitSignal, os.Interrupt)

	client.initialize()

	clientui := NewClientUI(client)
	// this runs ui event loop
	clientui.initUi()
}
