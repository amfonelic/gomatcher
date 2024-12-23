package endpoints

import (
	"log"
	"os"
	"regexp"
	"sync"

	"github.com/amfonelic/gomatcher/pkg/helpers"
)

type IEndpoint interface {
	HandleRequest(req any) error
	SetupServer()
	RunServer()
	String() string
	GetData() chan string
}

type DataManager struct {
	mu   sync.Mutex
	Data map[string]string
}

func RunServers(eps []IEndpoint) {
	for i, ep := range eps {
		log.Printf("[INFO] Setting up endpoint #%v\n", i)
		ep.SetupServer()
		go ep.RunServer()
	}
}

func ComposeMatchPrintData(eps []IEndpoint, pattern *regexp.Regexp) {
	for {
		manager := &DataManager{
			Data: make(map[string]string),
		}
		var wg sync.WaitGroup
		for _, ep := range eps {
			wg.Add(1)
			go ComposeData(ep, manager, &wg)
		}
		wg.Wait()
		log.Printf("[INFO] Data from all endpoints has been composed")

		valuesInSlice := helpers.MapToSlice(manager.Data)

		patterns, findErr := helpers.FindPatterns(pattern, valuesInSlice)
		if findErr != nil {
			log.Fatalf("[ERROR] %v", findErr)
		}
		isMatched, matchErr := helpers.AllStringsAreEqual(patterns)
		if matchErr != nil {
			log.Fatalf("[ERROR] %v", matchErr)
		}

		stdoutLogger := log.New(os.Stdout, "", log.LstdFlags)
		stdoutLogger.Println(manager.Data)
		stdoutLogger.Printf("Matched: %v", isMatched)

	}
}

func ComposeData(ch IEndpoint, manager *DataManager, wg *sync.WaitGroup) {
	defer wg.Done()
	if val, ok := <-ch.GetData(); ok {
		manager.mu.Lock()
		manager.Data[ch.String()] = val
		manager.mu.Unlock()
	}
}
