package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Log is the initialized logrus logger singleton for all of this project.
// Its output format is JSON and outputs to STDOUT.
var Log = logrus.New()
func init() {
	Log.Out = os.Stdout
	Log.SetFormatter(&logrus.JSONFormatter{})
}
