package GOEventBus

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
)

type Serializer struct {
	buffer bytes.Buffer
}

// Serializer constructor
func NewSerializer() *Serializer {
	return &Serializer{
		bytes.Buffer{},
	}
}

// Serialize serializes data into []byte
func (serializer *Serializer) Serialize(data map[string]any) []byte {
	enc := gob.NewEncoder(&serializer.buffer)
	err := enc.Encode(data)
	if err != nil {
		panic(err)
	}
	return serializer.buffer.Bytes()
}

// Deserialize deserializes data []byte to map[string]any
func (serializer *Serializer) Deserialize(data []byte) map[string]any {
	var result map[string]any
	if err := json.Unmarshal(data, &result); err != nil {
		panic(err)
	}
	return result
}
