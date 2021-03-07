package main

import (
	"github.com/phumberdroz/irsdk-go/pkg/sdk"
	"github.com/phumberdroz/irsdk-go/pkg/session"
	"github.com/phumberdroz/irsdk-go/pkg/telemetry"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
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
	logrus.Info(sessionData.WeekendInfo.TrackCity)
	tickSleepTime := time.Second / time.Duration(int(sdk.Header.TickRate))
	for {
		t1 := time.Now()
		values := sdk.ReadVariableValues()
		if values {
			var tData telemetry.TelemetryData
			err := sdk.Unmarshal(&tData)
			if err != nil {
				logrus.Error(err)
			}
			logrus.Infof("RPM: %f", tData.RPM)

			logrus.Infof("tick took %s", time.Since(t1))
		}
		time.Sleep(tickSleepTime)
	}

}
