package decoder

import "fmt"

type IDataDecoder interface {
	CheckFormat(input []byte) bool
	Decode(input []byte) string
}

func DetectFormat(input []byte) (IDataDecoder, error) {
	decoders := []IDataDecoder{
		JSONType(""),
		MsgPackType(""),
	}

	for _, d := range decoders {
		if d.CheckFormat(input) {
			switch d.(type) {
			case JSONType:
				return JSONType(""), nil
			case MsgPackType:
				return MsgPackType(""), nil
			}
		}
	}
	return nil, fmt.Errorf("data type not impemented")
}
