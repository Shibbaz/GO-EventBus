package pkg

import (
	"bytes"
	"encoding/json"
)

func Serialize(agregate Agregate) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(agregate)
	return b.Bytes(), err
}

func Deserialize(b []byte) (Agregate, error) {
	var msg Agregate
	buf := bytes.NewBuffer(b)
	decoder := json.NewDecoder(buf)
	err := decoder.Decode(&msg)
	return msg, err
}
