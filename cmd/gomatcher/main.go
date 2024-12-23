package main

import (
	"log"
	_ "net/http/pprof"

	"github.com/amfonelic/gomatcher/internal/config"
	"github.com/amfonelic/gomatcher/internal/endpoints"
)

func main() {
	eps, parseEpsErr := config.ParseEndpoints()
	if parseEpsErr != nil {
		log.Fatalf("[ERROR] Error while parsing ENV (endpoints). Error: %v", parseEpsErr)
	}
	pattern, parsePatternsErr := config.ParsePatterns()
	if parsePatternsErr != nil {
		log.Fatalf("[ERROR] Error while parsing ENV (patterns). Error: %v", parsePatternsErr)
	}

	endpoints.RunServers(eps)
	go endpoints.ComposeMatchPrintData(eps, pattern)

	select {}
}
