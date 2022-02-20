package GolangLogging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestFormatter(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Hello Logger")
	logger.Warn("Hello Logger")
	logger.Error("Hello Logger")
}
