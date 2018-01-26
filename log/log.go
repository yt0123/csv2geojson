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

func Init() {
	formatter := new(logrus.TextFormatter)
	formatter.FullTimestamp = true
	formatter.QuoteEmptyFields = true

	Logger.Formatter = formatter
	Logger.Out = os.Stdout
	Logger.SetLevel(logrus.DebugLevel)
}
