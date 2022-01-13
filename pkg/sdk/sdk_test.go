package sdk_test

import (
	"os"
	"testing"

	"github.com/0xkalle/irsdk-go/pkg/sdk"
	"github.com/0xkalle/irsdk-go/pkg/session"
	"github.com/0xkalle/irsdk-go/pkg/telemetry"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
)

func TestInit(t *testing.T) {
	file, err := os.Open("../../bmwm4gt3_bathurst 2021-02-12 19-18-01.ibt")
	require.NoError(t, err)
	init, err := sdk.InitSdk(file)
	require.NoError(t, err)
	require.True(t, init.SessionStatusOK())
	data, err := init.GetSessionData()
	require.NoError(t, err)
	var sessionData session.SessionData
	err = yaml.Unmarshal(data, &sessionData)
	require.NoError(t, err)
	assert.Equal(t, "Bathurst", sessionData.WeekendInfo.TrackCity)
	//fmt.Println("name,description,unit")
	//for _, v := range init.TelemetryVars.Vars {
	//	fmt.Printf("%s,%s,%s\n", v.Name, v.Desc, v.Unit)
	//}
	var tData telemetry.TelemetryData
	err = init.Unmarshal(&tData)
	assert.NoError(t, err)
	assert.Equal(t, "", tData.CarIdxRPM)
}
