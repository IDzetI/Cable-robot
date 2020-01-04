package utils

import (
	"encoding/json"
	"math"
	"unsafe"
)

func Uint16Float64(u []uint16) float64 {
	return math.Float64frombits(uint64(u[3])<<48 | uint64(u[2])<<32 | uint64(u[1])<<16 | uint64(u[0]))
}

func Float64Uint16(number float64) []uint16 {
	return (*[4]uint16)(unsafe.Pointer(&number))[:]
}

func ToString(i interface{}) string {
	b, err := json.Marshal(i)
	if err != nil {
		return err.Error()
	}
	return string(b)
}
