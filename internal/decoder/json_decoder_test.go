package decoder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONType_CheckFormat(t *testing.T) {
	validJSONData := []byte(`{"key": "value"}`)
	invalidJSONData := []byte("not json")

	jsonType := JSONType("")
	assert.True(t, jsonType.CheckFormat(validJSONData), "Expected valid JSON data to pass CheckFormat")
	assert.False(t, jsonType.CheckFormat(invalidJSONData), "Expected invalid JSON data to fail CheckFormat")
}

func TestJSONType_Decode(t *testing.T) {
	validJSONData := []byte(`{"key": "value"}`)

	jsonType := JSONType("")
	result := jsonType.Decode(validJSONData)
	assert.Equal(t, string(validJSONData), result, "Expected Decode to return the input JSON string unchanged")
}
