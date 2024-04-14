package GOEventBus

import (
	"encoding/json"
	"unsafe"
)

type Serializer struct{}

func (serializer *Serializer) Serialize(data map[string]interface{}) []byte {
	return *(*[]byte)(unsafe.Pointer(&data))
}

func (serializer *Serializer) Deserialize(data []byte) map[string]interface{} {
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		panic(err)
	}
	return result
}
