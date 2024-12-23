package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockEndpoint struct {
	mock.Mock
}

func (m *MockEndpoint) SomeEndpointMethod() {
}

func setupEnv(envVars map[string]string) {
	for key, value := range envVars {
		os.Setenv(key, value)
	}
}

func teardownEnv(envVars map[string]string) {
	for key := range envVars {
		os.Unsetenv(key)
	}
}

func TestParseEndpoints_HTTP_LessThanTwoEndpoints(t *testing.T) {
	envVars := map[string]string{
		"SERVER_TYPE":     "http",
		"HTTP_ENDPOINT_0": "/endpoint0",
		"HTTP_ENDPOINT_1": "",
	}

	setupEnv(envVars)
	defer teardownEnv(envVars)

	eps, err := ParseEndpoints()

	assert.Nil(t, eps)
	assert.EqualError(t, err, "less then two endpoints are passed")
}

func TestParseEndpoints_UnsupportedServerType(t *testing.T) {
	envVars := map[string]string{
		"SERVER_TYPE": "unsupported",
	}

	setupEnv(envVars)
	defer teardownEnv(envVars)

	eps, err := ParseEndpoints()

	assert.Nil(t, eps)
	assert.EqualError(t, err, "server type is not implemented")
}

func TestParseEndpoints_NoServerType(t *testing.T) {
	envVars := map[string]string{
		"SERVER_TYPE": "",
	}

	setupEnv(envVars)
	defer teardownEnv(envVars)

	eps, err := ParseEndpoints()

	assert.Nil(t, eps)
	assert.EqualError(t, err, "server type is not implemented")
}
