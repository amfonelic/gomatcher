package config

import (
	"fmt"
	"regexp"

	"github.com/amfonelic/gomatcher/pkg/env"
)

var predefinedPatterns = map[string]string{
	"uuid": `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}`,
}

func ParsePatterns() (*regexp.Regexp, error) {
	pattern := env.GetEnv[string]("PATTERN")
	if predefinedPattern, exists := predefinedPatterns[pattern]; exists {
		pattern = predefinedPattern
	}
	if pattern == "" {
		return nil, fmt.Errorf("pattern cannot be empty")
	}
	return regexp.Compile(pattern)
}
