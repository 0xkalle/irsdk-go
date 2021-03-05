package sdk

import (
	"fmt"
	"github.com/phumberdroz/irsdk-go/pkg/session"
	"golang.org/x/text/encoding/charmap"
	"io"
	"sync"
	"time"
)

type reader interface {
	io.Reader
	io.ReaderAt
	io.Closer
}

type SDK struct {
	r             reader
	h             Header
	s             session.SessionData
	TelemetryVars TelemetryVars
	lastValidData int64
	Timeout       time.Duration
}

type TelemetryVars struct {
	lastVersion int
	Vars        map[string]Variable
	mux         sync.Mutex
}

type Variable struct {
	VarType     int // irsdk_VarType
	Offset      int // Offset fron start of buffer row
	Count       int // number of entrys (array) so length in bytes would be irsdk_VarTypeBytes[type] * Count
	CountAsTime bool
	Name        string
	Desc        string
	Unit        string
	Value       interface{}
	rawBytes    []byte
}

func Init(r reader) (SDK, error) {
	if r == nil {
		return SDK{}, fmt.Errorf("not yet implemented")
	}
	/*
		Todo consider sending event
		https://github.com/vipoo/irsdk/blob/a14be906d1a6198aa0c6423022ed5246f20be3bb/irsdk_utils.cpp#L87-L101
	*/
	sdk := SDK{
		r:             r,
		lastValidData: 0,
		Timeout:       30 * time.Second,
	}
	h, err := readHeader(sdk.r)
	if err != nil {
		return SDK{}, err
	}
	sdk.h = h
	//if sdk.tVars != nil {
	//	sdk.tVars.vars = nil
	//}
	if !sdk.SessionStatusOK() {
		return SDK{}, fmt.Errorf("session status is not okay")
	}
	sdk.TelemetryVars = readVariableHeaders(sdk.r, &sdk.h)
	sdk.readVariableValues()
	return sdk, nil
}

func (s SDK) GetSessionData() ([]byte, error) {
	if !s.SessionStatusOK() {
		return nil, fmt.Errorf("session Status not ok")
	}
	rbuf := make([]byte, s.h.SessionInfoLen)
	_, err := s.r.ReadAt(rbuf, int64(s.h.SessionInfoOffset))
	if err != nil {
		return nil, err
	}

	dec := charmap.Windows1258.NewDecoder()
	rbuf, err = dec.Bytes(rbuf)
	if err != nil {
		return nil, err
	}

	return rbuf[:len(rbuf)-5], nil
}

func (s SDK) SessionStatusOK() bool {
	return s.h.Status > 0
}
