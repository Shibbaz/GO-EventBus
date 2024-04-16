package GOEventBus

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

type Serializer struct {
	buffer bytes.Buffer
}

func NewSerializer() *Serializer {
	return &Serializer{
		bytes.Buffer{},
	}
}

func (serializer *Serializer) Serialize(data map[string]any) []byte {
	enc := gob.NewEncoder(&serializer.buffer)
	err := enc.Encode(data)
	if err != nil {
		panic(err)
	}
	return serializer.buffer.Bytes()
}

func (serializer *Serializer) Deserialize(data []byte) map[string]any {
	var result map[string]any
	if err := json.Unmarshal(data, &result); err != nil {
		panic(err)
	}
	return result
}
