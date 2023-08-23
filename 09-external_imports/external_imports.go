package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	// logrus.SetLevel(logrus.DebugLevel)
	logrus.SetLevel(logrus.WarnLevel)

	logrus.Debug("Debug Message")
	logrus.Warning("Debug Message Warning")
}
