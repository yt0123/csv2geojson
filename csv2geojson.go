package csv2geojson

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"unicode/utf8"

	"github.com/ty-edelweiss/csv2geojson/geo"
	"github.com/ty-edelweiss/csv2geojson/log"
)

var (
	report = log.NewReport()
	logger = log.AppLogger.SetName("csv2geojson")
)

type Converter struct {
	FileName     string
	BufferLength int64
	Reader       *csv.Reader
	Options      *Options
}

func NewConverter(inputPath string, options *Options) Converter {
	fp, err := os.Open(inputPath)
	if err != nil {
		logger.Fatal(err)
	}

	fi, err := fp.Stat()
	if err != nil {
		logger.Fatal(err)
	}
	logger.WithField("filename", fi.Name()).WithField("length", fi.Size()).Info("Target file is the following.")

	delimiter, size := utf8.DecodeRuneInString(options.Delimiter)
	if size != 1 {
		logger.Fatal(errors.New("Delimiter size is invalid"))
	}
	logger.WithField("delimiter", delimiter).Info("Delimiter for csv file set the following.")

	reader := csv.NewReader(fp)
	reader.Comma = delimiter
	reader.LazyQuotes = !options.Quotes

	logger.WithField("option", *options).Debug("Application options from command line.")

	return Converter{fi.Name(), fi.Size(), reader, options}
}

func (c *Converter) Do() {
	report.FormatCMessage("%s is importing", c.FileName)

	records, err := c.Reader.ReadAll()
	if err != nil {
		logger.Fatal(err)
	}
	logger.WithField("headers", records[0]).Debug("Input csv file headers is following.")

	report.PrintMessage("ok")

	report.FormatMessage("converting %s [ %d bytes ] :", c.FileName, int(c.BufferLength))

	buf, err := geo.Build(c.Options.Type, c.Options.Longitude, c.Options.Latitude, c.Options.Key, records[0], records[1:], c.Options.Limit)
	if err != nil {
		logger.Fatal(err)
	}
	logger.WithField("data", string(buf)).Debug("Geojson conversion done. and result is following.")

	if c.Options.Preformat {
		var out bytes.Buffer
		json.Indent(&out, buf, "", "\t")
		buf = out.Bytes()
	}

	deploy(c.Options.Output, buf)
}

func deploy(path string, content []byte) {
	cwd, err := os.Getwd()
	if err != nil {
		logger.Fatal(err)
	}

	output := cwd + "/result.geojson"
	if path != "" {
		output = path
	}

	report.CMessage("deploying")

	ioutil.WriteFile(output, content, os.ModePerm)

	report.PrintMessage("done")
}
