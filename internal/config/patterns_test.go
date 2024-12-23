package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsePatterns_PredefinedPattern(t *testing.T) {
	envVars := map[string]string{
		"PATTERN": "uuid",
	}
	setupEnv(envVars)
	defer teardownEnv(envVars)

	regex, err := ParsePatterns()

	assert.NoError(t, err)
	assert.NotNil(t, regex)
	assert.Equal(t, predefinedPatterns["uuid"], regex.String())
}

func TestParsePatterns_CustomPattern(t *testing.T) {
	envVars := map[string]string{
		"PATTERN": `^\d{3}-\d{2}-\d{4}$`,
	}
	setupEnv(envVars)
	defer teardownEnv(envVars)

	regex, err := ParsePatterns()

	assert.NoError(t, err)
	assert.NotNil(t, regex)
	assert.Equal(t, `^\d{3}-\d{2}-\d{4}$`, regex.String())
}

func TestParsePatterns_InvalidCustomPattern(t *testing.T) {
	envVars := map[string]string{
		"PATTERN": `(*`,
	}
	setupEnv(envVars)
	defer teardownEnv(envVars)

	regex, err := ParsePatterns()

	assert.Error(t, err)
	assert.Nil(t, regex)
}

func TestParsePatterns_EmptyPattern(t *testing.T) {
	envVars := map[string]string{
		"PATTERN": "",
	}
	setupEnv(envVars)
	defer teardownEnv(envVars)

	regex, err := ParsePatterns()

	assert.Error(t, err)
	assert.Nil(t, regex)
}
