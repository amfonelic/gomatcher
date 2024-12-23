package decoder

import (
	"encoding/json"
	"fmt"

	"github.com/vmihailenco/msgpack"
)

type MsgPackType string

func (mt MsgPackType) CheckFormat(input []byte) bool {
	var data interface{}
	if err := msgpack.Unmarshal(input, &data); err != nil {
		return false
	}
	return true
}

func (mt MsgPackType) Decode(input []byte) string {
	var data interface{}
	msgpack.Unmarshal(input, &data)
	switch data := data.(type) {
	case string:
		return data
	case map[string]interface{}:
		jsonBytes, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			return ""
		}
		return string(jsonBytes)
	default: //int, float, etc...
		return fmt.Sprintf("%v", data)
	}
}
