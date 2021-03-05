package main

import (
	"github.com/phumberdroz/irsdk-go/pkg/sdk"
	"github.com/phumberdroz/irsdk-go/pkg/session"
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
			getVar, err := sdk.GetVar("RPM")
			if err != nil {
				logrus.WithError(err).Error("failed to get rpm")
				continue
			}
			logrus.Info("rpm is ", getVar.Value)
			logrus.Infof("Loop took %d", time.Now().Sub(t1).Nanoseconds())
		}
		time.Sleep(tickSleepTime)
	}

}
