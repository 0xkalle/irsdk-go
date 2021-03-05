package sdk

import (
	"fmt"
	"log"
	"time"

	"github.com/phumberdroz/irsdk-go/pkg/helpers"
)

func readVariableHeaders(r reader, h *Header) TelemetryVars {
	vars := TelemetryVars{Vars: make(map[string]Variable, h.NumVars)}
	for i := 0; i < int(h.NumVars); i++ {
		rbuf := make([]byte, 144)
		_, err := r.ReadAt(rbuf, int64(int(h.HeaderOffset)+i*144))
		if err != nil {
			log.Fatal(err)
		}

		v := Variable{
			VarType:     helpers.Byte4ToInt(rbuf[0:4]),
			Offset:      helpers.Byte4ToInt(rbuf[4:8]),
			Count:       helpers.Byte4ToInt(rbuf[8:12]),
			CountAsTime: int(rbuf[12]) > 0,
			Name:        helpers.BytesToString(rbuf[16:48]),
			Desc:        helpers.BytesToString(rbuf[48:112]),
			Unit:        helpers.BytesToString(rbuf[112:144]),
		}
		//log.Printf("%s,%s,%s", v.Name, v.Desc, v.Unit)
		vars.Vars[v.Name] = v
	}
	return vars
}

type varBuffer struct {
	tickCount int // used to detect changes in data
	bufOffset int // offset from header
}

func (s SDK) FindLatestBuffer() varBuffer {
	var vb varBuffer
	foundTickCount := 0
	for i := 0; i < int(s.Header.NumBuf); i++ {
		rbuf := make([]byte, 16)
		_, err := s.r.ReadAt(rbuf, int64(48+i*16))
		if err != nil {
			log.Fatal(err)
		}
		currentVb := varBuffer{
			tickCount: helpers.Byte4ToInt(rbuf[0:4]),
			bufOffset: helpers.Byte4ToInt(rbuf[4:8]),
		}
		//fmt.Printf("BUFF?: %+v\n", currentVb)
		if foundTickCount < currentVb.tickCount {
			foundTickCount = currentVb.tickCount
			vb = currentVb
		}
	}
	//fmt.Printf("BUFF: %+v\n", vb)
	return vb
}

func (sdk *SDK) ReadVariableValues() bool {
	newData := false
	if sdk.SessionStatusOK() {
		// find latest buffer for variables
		vb := sdk.FindLatestBuffer()
		sdk.TelemetryVars.mux.Lock()
		defer sdk.TelemetryVars.mux.Unlock()
		if sdk.TelemetryVars.lastVersion < vb.tickCount {
			newData = true
			sdk.TelemetryVars.lastVersion = vb.tickCount
			sdk.lastValidData = time.Now().Unix()
			for varName, v := range sdk.TelemetryVars.Vars {
				var rbuf []byte
				switch v.VarType {
				case 0:
					rbuf = make([]byte, 1)
					_, err := sdk.r.ReadAt(rbuf, int64(vb.bufOffset+v.Offset))
					if err != nil {
						log.Fatal(err)
					}
					v.Value = string(rbuf[0])
				case 1:
					rbuf = make([]byte, 1)
					_, err := sdk.r.ReadAt(rbuf, int64(vb.bufOffset+v.Offset))
					if err != nil {
						log.Fatal(err)
					}
					v.Value = int(rbuf[0]) > 0
				case 2:
					rbuf = make([]byte, 4)
					_, err := sdk.r.ReadAt(rbuf, int64(vb.bufOffset+v.Offset))
					if err != nil {
						log.Fatal(err)
					}
					v.Value = helpers.Byte4ToInt(rbuf)
				case 3:
					rbuf = make([]byte, 4)
					_, err := sdk.r.ReadAt(rbuf, int64(vb.bufOffset+v.Offset))
					if err != nil {
						log.Fatal(err)
					}
					v.Value = helpers.Byte4toBitField(rbuf)
					fmt.Println(v.Value)
				case 4:
					rbuf = make([]byte, 4)
					_, err := sdk.r.ReadAt(rbuf, int64(vb.bufOffset+v.Offset))
					if err != nil {
						log.Fatal(err)
					}
					v.Value = helpers.Byte4ToFloat(rbuf)
				case 5:
					rbuf = make([]byte, 8)
					_, err := sdk.r.ReadAt(rbuf, int64(vb.bufOffset+v.Offset))
					if err != nil {
						log.Fatal(err)
					}
					v.Value = helpers.Byte8ToFloat(rbuf)
				}
				v.rawBytes = rbuf
				sdk.TelemetryVars.Vars[varName] = v
			}
		}
	}

	return newData
}

func (sdk *SDK) GetVar(name string) (Variable, error) {
	if !sdk.SessionStatusOK() {
		return Variable{}, fmt.Errorf("Session is not active")
	}
	sdk.TelemetryVars.mux.Lock()
	defer sdk.TelemetryVars.mux.Unlock()
	if v, ok := sdk.TelemetryVars.Vars[name]; ok {
		return v, nil
	}
	return Variable{}, nil
}
