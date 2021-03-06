package helpers

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

func Byte4ToInt(in []byte) (interface{}, error) {
	r := bytes.NewReader(in)
	if len(in) == 4 {
		i := int32(0)
		err := binary.Read(r, binary.LittleEndian, &i)
		if err != nil {
			return nil, err
		}
		return i, nil
	} else {
		i := make([]int32, len(in)/4)
		err := binary.Read(r, binary.LittleEndian, &i)
		if err != nil {
			return nil, err
		}
		return i, nil
	}
	//return nil, nil
}
func Byte4ToFloat(in []byte) (interface{}, error) {
	r := bytes.NewReader(in)
	if len(in) == 4 {
		i := float32(0)
		err := binary.Read(r, binary.LittleEndian, &i)
		if err != nil {
			return nil, err
		}
		return i, nil
	} else {
		i := make([]float32, len(in)/4)
		err := binary.Read(r, binary.LittleEndian, &i)
		if err != nil {
			return nil, err
		}
		return i, nil
	}
}
func Byte8ToFloat(in []byte) (interface{}, error) {
	r := bytes.NewReader(in)
	if len(in) == 8 {
		i := float64(0)
		err := binary.Read(r, binary.LittleEndian, &i)
		if err != nil {
			return nil, err
		}
		return i, nil
	} else {
		i := make([]float64, len(in)/8)
		err := binary.Read(r, binary.LittleEndian, &i)
		if err != nil {
			return nil, err
		}
		return i, nil
	}
}
func Byte4toBitField(in []byte) string {
	return fmt.Sprintf("0x%x", int(binary.LittleEndian.Uint32(in)))
}
func BytesToString(in []byte) string {
	return strings.TrimRight(string(in), "\x00")
}
