package log

import "github.com/sirupsen/logrus"

func Logrus() {
	logrus.Infof("/sirupsen/logrus info: %s", "hello logrus")
}