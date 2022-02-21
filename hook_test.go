package GolangLogging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
)

type SampleHook struct {
}

func (s SampleHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.WarnLevel,
		logrus.ErrorLevel,
	}
}

func (s SampleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("SampleHook", entry.Level, entry.Message)
	return nil
}

func TestHook(t *testing.T) {
	logger := logrus.New()
	logger.AddHook(&SampleHook{})
	logrus.SetLevel(logrus.TraceLevel)

	logger.Info("Hello Info")
	logger.Warn("Hello Info")
	logger.Error("Hello Info")
	logger.Debug("Hello Info")
}
