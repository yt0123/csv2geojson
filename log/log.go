package log

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	LIMITS    = 5
	PREFIX    = "application-csv2geojson-"
	EXTENSION = ".log"
	DIRECTORY = "tmp"
)

var (
	AppLogger = Logger{}
)

func init() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	path := filepath.Join(wd, DIRECTORY)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(DIRECTORY, os.ModePerm); err != nil {
			panic(err)
		}
	} else {
		if info, err := ioutil.ReadDir(path); err != nil {
			panic(err)
		} else if len(info) >= LIMITS {
			target := filepath.Join(path, info[0].Name())
			if err := os.Remove(target); err != nil {
				panic(err)
			}
		}
	}

	path = filepath.Join(path, PREFIX+time.Now().Format("20060102150405")+EXTENSION)
	fp, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}

	Init(fp, "app", false)
}

type Logger struct {
	Name string
	logrus.Logger
}

func Init(out *os.File, name string, verbose bool) {
	AppLogger.Name = name

	AppLogger.Out = out

	formatter := new(logrus.TextFormatter)
	formatter.FullTimestamp = true
	formatter.DisableColors = true
	formatter.QuoteEmptyFields = true

	AppLogger.Formatter = formatter
	if verbose {
		AppLogger.SetLevel(logrus.DebugLevel)
	} else {
		AppLogger.SetLevel(logrus.WarnLevel)
	}
}

func (l *Logger) SetName(name string) *Logger {
	l.Name = name
	return l
}

func (l *Logger) SetVerbose(verbose bool) *Logger {
	if verbose {
		AppLogger.SetLevel(logrus.DebugLevel)
	} else {
		AppLogger.SetLevel(logrus.WarnLevel)
	}
	return l
}
