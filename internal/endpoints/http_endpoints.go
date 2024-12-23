package endpoints

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/amfonelic/gomatcher/internal/decoder"
	"github.com/amfonelic/gomatcher/pkg/env"
)

var serverOnce sync.Once

type HTTPEndpoint struct {
	Path string
	Data chan string
}

type HTTPRequestWrapper struct {
	W http.ResponseWriter
	R *http.Request
}

func CreateHTTPEndpoint(path string) IEndpoint {
	return &HTTPEndpoint{Path: path, Data: make(chan string)}
}

func (he *HTTPEndpoint) HandleRequest(req any) error {
	httpReq, ok := req.(*HTTPRequestWrapper)
	if !ok {
		return fmt.Errorf("invalid request type: expected HTTPRequestWrapper, got %T", httpReq)
	}

	bodyBytes, bodyErr := io.ReadAll(httpReq.R.Body)
	if bodyErr != nil {
		httpReq.W.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(httpReq.W, "Error reading body: %v\n", bodyErr)
		return fmt.Errorf("error reading body")
	}

	dec, detectDecodeErr := decoder.DetectFormat(bodyBytes)
	if detectDecodeErr != nil {
		httpReq.W.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(httpReq.W, "Error detecting body format: %v\n", detectDecodeErr)
		return fmt.Errorf("error detecting body")
	}

	he.Data <- dec.Decode(bodyBytes)
	log.Printf("[INFO] Data has been given to HTTP endpoint with path: %v;", he.Path)

	return nil
}

func (he *HTTPEndpoint) SetupServer() {
	log.Printf("[INFO] Setting up endpoint with path: %v\n", he.Path)
	http.HandleFunc(he.Path, func(w http.ResponseWriter, r *http.Request) {
		if err := he.HandleRequest(&HTTPRequestWrapper{W: w, R: r}); err != nil {
			log.Printf("[ERROR] error while handling request: %v", err)
		}
	})
}

func (he *HTTPEndpoint) RunServer() {
	serverOnce.Do(func() {
		addr := fmt.Sprintf("%v:%d", env.GetEnv[string]("HTTP_SERVER_HOST"), env.GetEnv[int]("HTTP_SERVER_PORT"))
		log.Printf("[INFO] Starting HTTP Web Server. Addr: %v\n", addr)

		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Fatalf("error starting server: %v", err)
		}
	})
}

func (he *HTTPEndpoint) GetData() chan string {
	return he.Data
}

func (he *HTTPEndpoint) String() string {
	return fmt.Sprintf("HTTP (%v)", he.Path)
}
