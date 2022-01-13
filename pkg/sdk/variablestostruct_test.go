package sdk_test

import (
	"encoding/json"
	"fmt"
	"github.com/0xkalle/irsdk-go/pkg/sdk"
	"github.com/0xkalle/irsdk-go/pkg/telemetry"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestVariableToStruct(t *testing.T) {
	file, err := os.Open("../../bmwm4gt3_bathurst 2021-02-12 19-18-01.ibt")
	require.NoError(t, err)
	init, err := sdk.InitSdk(file)
	require.NoError(t, err)
	require.True(t, init.SessionStatusOK())
	var tData telemetry.TelemetryData
	err = init.Unmarshal(&tData)
	assert.NoError(t, err)
	fmt.Println(tData.CarIdxBestLapTime)
	marshal, err := json.Marshal(tData)
	require.NoError(t, err)
	fmt.Println(string(marshal))
}
