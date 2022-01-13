package main

import (
	"github.com/0xkalle/irsdk-go/pkg/sdk"
	"github.com/0xkalle/irsdk-go/pkg/session"
	"github.com/0xkalle/irsdk-go/pkg/telemetry"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"time"
)

func main() {
	sdk, err := sdk.InitSdk(nil)
	if err != nil {
		logrus.WithError(err).Error("failed to init")
		return
	}

	data, err := sdk.GetSessionData()
	if err != nil {
		logrus.WithError(err).Error("failed to get session data")
		return
	}
	var sessionData session.SessionData
	err = yaml.Unmarshal(data, &sessionData)
	if err != nil {
		logrus.WithError(err).Error("failed to unmarshal sessiondata")
		return
	}
	logrus.Info(sessionData.WeekendInfo.TrackName)
	tickSleepTime := time.Second / time.Duration(int(sdk.Header.TickRate))
	t1 := time.Now()
	for {
		values := sdk.ReadVariableValues()
		if values {
			var tData telemetry.TelemetryData
			err := sdk.Unmarshal(&tData)
			if err != nil {
				logrus.Error(err)
			}
			logrus.Infof("G: %t", tData.IsInGarage)
			logrus.Infof("T: %t", tData.IsOnTrack)
			logrus.Infof("TC: %t", tData.IsOnTrackCar)

			logrus.Infof("tick took %s", time.Since(t1))
			t1 = time.Now()
		}
		time.Sleep(tickSleepTime)
	}

}
