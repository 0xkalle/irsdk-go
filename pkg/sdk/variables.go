package sdk

import (
	"fmt"
	"log"
	"time"

	"github.com/0xkalle/irsdk-go/pkg/helpers"
)

// Todo make this return error
func readVariableHeaders(r reader, h *Header) TelemetryVars {
	vars := TelemetryVars{Vars: make(map[string]Variable, h.NumVars)}
	for i := 0; i < int(h.NumVars); i++ {
		rbuf := make([]byte, 144)
		_, err := r.ReadAt(rbuf, int64(int(h.HeaderOffset)+i*144))
		if err != nil {
			log.Fatal(err)
		}

		v := Variable{
			CountAsTime: int(rbuf[12]) > 0,
			Name:        helpers.BytesToString(rbuf[16:48]),
			Desc:        helpers.BytesToString(rbuf[48:112]),
			Unit:        helpers.BytesToString(rbuf[112:144]),
		}
		varType, err := helpers.Byte4ToInt(rbuf[0:4])
		if err != nil {
			log.Fatal(err)
		}
		v.VarType = varType.(int32)
		offSet, err := helpers.Byte4ToInt(rbuf[4:8])
		if err != nil {
			log.Fatal(err)
		}
		v.Offset = offSet.(int32)
		count, err := helpers.Byte4ToInt(rbuf[8:12])
		if err != nil {
			log.Fatal(err)
		}
		v.Count = count.(int32)
		//log.Printf("%s,%s,%s", v.Name, v.Desc, v.Unit)
		vars.Vars[v.Name] = v
	}
	return vars
}

type varBuffer struct {
	tickCount int32 // used to detect changes in data
	bufOffset int32 // offset from header
}

func (s SDK) FindLatestBuffer() varBuffer {
	var vb varBuffer
	foundTickCount := int32(0)
	for i := 0; i < int(s.Header.NumBuf); i++ {
		rbuf := make([]byte, 16)
		_, err := s.r.ReadAt(rbuf, int64(48+i*16))
		if err != nil {
			log.Fatal(err)
		}
		currentVb := varBuffer{}
		tickCount, err := helpers.Byte4ToInt(rbuf[0:4])
		if err != nil {
			log.Fatal(err)
		}
		currentVb.tickCount = tickCount.(int32)
		buffOffset, err := helpers.Byte4ToInt(rbuf[4:8])
		if err != nil {
			log.Fatal(err)
		}
		currentVb.bufOffset = buffOffset.(int32)

		if foundTickCount < currentVb.tickCount {
			foundTickCount = currentVb.tickCount
			vb = currentVb
		}
	}
	return vb
}

const (
	CharType     int32 = iota
	BoolType     int32 = iota
	IntType      int32 = iota
	BitfieldType int32 = iota
	FloatType    int32 = iota
	DoubleType   int32 = iota
)

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
				case CharType:
					rbuf = make([]byte, 1)
					_, err := sdk.r.ReadAt(rbuf, int64(vb.bufOffset+v.Offset))
					if err != nil {
						log.Fatal(err)
					}
					v.Value = string(rbuf[0])
				case BoolType:
					rbuf = make([]byte, 1*v.Count)
					_, err := sdk.r.ReadAt(rbuf, int64(vb.bufOffset+v.Offset))
					if err != nil {
						log.Fatal(err)
					}
					v.Value, err = helpers.ByteToBool(rbuf)
					if err != nil {
						log.Fatalln(err)
					}
				case IntType:
					rbuf = make([]byte, 4*v.Count)
					_, err := sdk.r.ReadAt(rbuf, int64(vb.bufOffset+v.Offset))
					if err != nil {
						log.Fatal(err)
					}

					v.Value, err = helpers.Byte4ToInt(rbuf)
					if err != nil {
						log.Fatalln(err)
					}
				case BitfieldType:
					rbuf = make([]byte, 4)
					_, err := sdk.r.ReadAt(rbuf, int64(vb.bufOffset+v.Offset))
					if err != nil {
						log.Fatal(err)
					}
					v.Value = helpers.Byte4toBitField(rbuf)
				case FloatType:
					rbuf = make([]byte, 4*v.Count)
					_, err := sdk.r.ReadAt(rbuf, int64(vb.bufOffset+v.Offset))
					if err != nil {
						log.Fatal(err)
					}
					v.Value, _ = helpers.Byte4ToFloat(rbuf)
				case DoubleType:
					rbuf = make([]byte, 8*v.Count)
					_, err := sdk.r.ReadAt(rbuf, int64(vb.bufOffset+v.Offset))
					if err != nil {
						log.Fatal(err)
					}
					v.Value, _ = helpers.Byte8ToFloat(rbuf)
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
