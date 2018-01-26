package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

const (
	file = "./csv2geojson.log"
)

var (
	Logger = logrus.New()
)

func Init(verbose bool) {
	formatter := new(logrus.TextFormatter)
	formatter.FullTimestamp = true
	formatter.QuoteEmptyFields = true

	Logger.Formatter = formatter
	Logger.Out = os.Stdout
	if verbose {
		Logger.SetLevel(logrus.DebugLevel)
	} else {
		Logger.SetLevel(logrus.WarnLevel)
	}
}
