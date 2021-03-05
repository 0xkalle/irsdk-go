package sdk_test

import (
	"github.com/phumberdroz/irsdk-go/pkg/sdk"
	"github.com/phumberdroz/irsdk-go/pkg/session"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"
	"os"
	"testing"
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
	//for i := 0; i < 31015; i++ {
	//	init.ReadVariableValues()
	//	getVar, err := init.GetVar("RPM")
	//	require.NoError(t, err)
	//	fmt.Println(getVar.Value)
	//}
}
