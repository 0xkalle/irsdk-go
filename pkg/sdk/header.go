package sdk

import (
	"bytes"
	"encoding/binary"
)

// Header please keep in mind that order is important
type Header struct {
	Version  int32 // api version 1 for now
	Status   int32 // bitfield using StatusField
	TickRate int32 // ticks per second (60 or 360 etc)

	// session information, updated periodicaly
	SessionInfoUpdate int32 // Incremented when session info changes
	SessionInfoLen    int32 // Length in bytes of session info string
	SessionInfoOffset int32 // Session info, encoded in YAML format

	// State data, output at tickRate
	NumVars      int32 // length of array pointed to by varHeaderOffset
	HeaderOffset int32 // offset to VarHeader[numVars] array, Describes the variables recieved in varBuf

	NumBuf int32
	BufLen int32
}

func readHeader(r reader) (Header, error) {
	h := Header{}
	rbuf := make([]byte, binary.Size(h))
	_, err := r.ReadAt(rbuf, 0)
	if err != nil {
		return Header{}, err
	}

	err = binary.Read(bytes.NewReader(rbuf), binary.LittleEndian, &h)
	if err != nil {
		return Header{}, err
	}
	return h, nil
}
