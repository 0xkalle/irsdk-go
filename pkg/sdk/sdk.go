package sdk

import (
	"fmt"
	"io"
	"sync"
	"time"

	"bitbucket.org/avd/go-ipc/mmf"
	"github.com/phumberdroz/irsdk-go/internal/shm"
	"github.com/phumberdroz/irsdk-go/pkg/session"
	"golang.org/x/text/encoding/charmap"
)

type reader interface {
	io.Reader
	io.ReaderAt
	//io.Closer
}

type SDK struct {
	r             reader
	Header        Header
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

const fileMapName string = "Local\\IRSDKMemMapFileName"
const fileMapSize int = 1164 * 1024

func InitSdk(r reader) (SDK, error) {
	if r == nil {
		object, err := shm.NewWindowsNativeMemoryObject(fileMapName, 0, fileMapSize)
		if err != nil {
			return SDK{}, err
		}
		roRegion, err := mmf.NewMemoryRegion(object, mmf.MEM_READ_ONLY, 0, fileMapSize)
		if err != nil {
			return SDK{}, err
		}
		regionReader := mmf.NewMemoryRegionReader(roRegion)
		r = regionReader
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
	sdk.Header = h
	//if sdk.tVars != nil {
	//	sdk.tVars.vars = nil
	//}
	if !sdk.SessionStatusOK() {
		return SDK{}, fmt.Errorf("session status is not okay")
	}
	sdk.TelemetryVars = readVariableHeaders(sdk.r, &sdk.Header)
	sdk.ReadVariableValues()
	return sdk, nil
}

func (s SDK) GetSessionData() ([]byte, error) {
	if !s.SessionStatusOK() {
		return nil, fmt.Errorf("session Status not ok")
	}
	rbuf := make([]byte, s.Header.SessionInfoLen)
	_, err := s.r.ReadAt(rbuf, int64(s.Header.SessionInfoOffset))
	if err != nil {
		return nil, err
	}

	dec := charmap.Windows1258.NewDecoder()
	rbuf, err = dec.Bytes(rbuf)
	if err != nil {
		return nil, err
	}
	var r []byte
	for _, b := range rbuf {
		if b != byte(0) {
			r = append(r, b)
		}
	}
	return r, nil
	//return rbuf[:len(rbuf)], nil
}

func (s SDK) SessionStatusOK() bool {
	return s.Header.Status > 0
}
