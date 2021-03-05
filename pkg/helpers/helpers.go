package helpers

import (
	"encoding/binary"
	"fmt"
	"math"
	"strings"
)

func Byte4ToInt(in []byte) int {
	return int(binary.LittleEndian.Uint32(in))
}
func Byte4ToFloat(in []byte) float32 {
	bits := binary.LittleEndian.Uint32(in)
	return math.Float32frombits(bits)
}
func Byte8ToFloat(in []byte) float64 {
	bits := binary.LittleEndian.Uint64(in)
	return math.Float64frombits(bits)
}
func Byte4toBitField(in []byte) string {
	return fmt.Sprintf("0x%x", int(binary.LittleEndian.Uint32(in)))
}
func BytesToString(in []byte) string {
	return strings.TrimRight(string(in), "\x00")
}
