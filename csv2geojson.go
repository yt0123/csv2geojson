package csv2geojson

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"unicode/utf8"

	"github.com/ty-edelweiss/csv2geojson/geo"
	"github.com/ty-edelweiss/csv2geojson/log"
)

type Converter struct {
	Reader  *csv.Reader
	Options *Options
}

func NewConverter(inputPath string, options *Options) Converter {
	fp, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}

	delimiter, size := utf8.DecodeRuneInString(options.Delimiter)
	if size != 1 {
		panic(errors.New("Delimiter size is invalid"))
	}
	log.Logger.WithField("delimiter", delimiter).Info("Delimiter for csv file set the following.")

	reader := csv.NewReader(fp)
	reader.Comma = delimiter
	reader.LazyQuotes = !options.Quotes

	log.Logger.WithField("option", *options).Debug("Application options from command line.")

	return Converter{reader, options}
}

func (c *Converter) Do() {
	var headers []string
	var records [][]string

	index := 0
	for {
		record, err := c.Reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Logger.Fatal(err)
		}

		if index == 0 {
			headers = record
		} else {
			records = append(records, record)
		}

		index = index + 1
	}
	log.Logger.WithField("headers", headers).Debug("Input csv file headers is following.")

	buf, err := geo.Build(c.Options.Type, c.Options.Longitude, c.Options.Latitude, c.Options.Key, headers, records)
	if err != nil {
		log.Logger.Fatal(err)
	}
	log.Logger.WithField("data", string(buf)).Info("Geojson conversion done. and result is following.")

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
		panic(err)
	}

	output := cwd + "/result.geojson"
	if path != "" {
		output = path
	}

	ioutil.WriteFile(output, content, os.ModePerm)
}
