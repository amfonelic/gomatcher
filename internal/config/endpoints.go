package config

import (
	"fmt"
	"log"

	"github.com/amfonelic/gomatcher/internal/endpoints"
	"github.com/amfonelic/gomatcher/pkg/env"
)

type serverType string

const (
	HTTP serverType = "http"
)

func ParseEndpoints() ([]endpoints.IEndpoint, error) {
	var eps []endpoints.IEndpoint
	switch env.GetEnv[string]("SERVER_TYPE") {
	case "http":
		for i := 0; ; i++ {
			endpoint := fmt.Sprintf("HTTP_ENDPOINT_%v", i)
			path := env.GetEnv[string](endpoint)
			if path == "" {
				break
			}
			log.Printf("[INFO] Parse endpoint %v with path %v\n", endpoint, path)
			eps = append(eps, endpoints.CreateHTTPEndpoint(path))
		}

		if len(eps) < 2 {
			return nil, fmt.Errorf("less then two endpoints are passed")
		}
	default:
		return nil, fmt.Errorf("server type is not implemented")
	}
	return eps, nil
}
