package util

import (
	"bytes"
	"encoding/binary"
	"math"
)

func BytesToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

func Float64ToBytes(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

func BytesToStruct(b []byte) (i interface{}, err error) {
	err = binary.Read(bytes.NewReader(b), binary.BigEndian, &i)
	return
}

func StructToBytes(obj interface{}) (b []byte, err error) {
	buf := &bytes.Buffer{}
	err = binary.Write(buf, binary.BigEndian, obj)
	if err != nil {
		return
	}
	b = buf.Bytes()
	return
}
