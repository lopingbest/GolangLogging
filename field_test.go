package GolangLogging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "galih").Info("Hello World")

	logger.WithField("username", "setiadi").
		WithField("name", "Galih Setiadi").
		Info("Hello World")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "galih",
		"name":     "Galih Setiadi",
	}).Info("Hello World")
}
