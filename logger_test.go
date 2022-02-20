package GolangLogging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLogger(t *testing.T) {
	logger := logrus.New()

	logger.Println("Logger")
	fmt.Println("Hello logger")
}
