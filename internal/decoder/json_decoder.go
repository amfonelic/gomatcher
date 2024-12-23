package decoder

import (
	"encoding/json"
)

type JSONType string

func (jt JSONType) CheckFormat(input []byte) bool {
	var data interface{}
	if err := json.Unmarshal(input, &data); err != nil {
		return false
	}
	return true
}

func (jt JSONType) Decode(input []byte) string {
	return string(input)
}
