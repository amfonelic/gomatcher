package helpers

import (
	"fmt"
	"regexp"
)

func MapToSlice(m map[string]string) []string {
	newSlice := make([]string, 0, len(m))
	for _, v := range m {
		newSlice = append(newSlice, v)
	}
	return newSlice
}

func SlicesToMap(keys, values []string) map[string]string {
	newMap := make(map[string]string)
	for i, v := range keys {
		newMap[v] = values[i]
	}
	return newMap
}

func AllStringsAreEqual(slice []string) (bool, error) {
	if len(slice) <= 1 {
		return false, fmt.Errorf("len(slice) <= 1")
	}

	for _, str := range slice[1:] {
		if str != slice[0] {
			return false, nil
		}
	}

	return true, nil
}

func FindPatterns(pattern *regexp.Regexp, slice []string) ([]string, error) {
	patterns := make([]string, len(slice))
	for i, str := range slice {
		match := pattern.FindString(str)
		if match == "" {
			return nil, fmt.Errorf("cannot find pattern")
		}
		patterns[i] = match
	}
	return patterns, nil
}
