package main

import (
	"github.com/trasa/watchmud/message"
	"log"
)

func main() {
	req := message.LookRequest{}
	resp := message.LookResponse{}
	log.Printf("req: %s resp: %s", req, resp)
}
