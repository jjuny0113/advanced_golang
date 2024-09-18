package utils

import (
	"bytes"
	"encoding/gob"
)

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func ToBytes(i interface{}) []byte {
	var aBuffer bytes.Buffer
	encoder := gob.NewEncoder(&aBuffer)
	HandleError(encoder.Encode(i))
	return aBuffer.Bytes()
}

func FromBytes(i interface{}, data []byte) {
	encoder := gob.NewDecoder(bytes.NewReader(data))
	HandleError(encoder.Decode(i))
}
