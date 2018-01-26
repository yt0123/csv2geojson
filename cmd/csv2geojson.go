package main

import (
	"github.com/ty-edelweiss/csv2geojson"
	"github.com/ty-edelweiss/csv2geojson/log"
)

var (
	logger = log.AppLogger.SetName("main")
)

func main() {
	var opts csv2geojson.Options
	args := opts.Parse()

	logger.SetVerbose(opts.Verbose)
	logger.Info("Application starting on command line.")

	converter := csv2geojson.NewConverter(args[0], &opts)

	converter.Do()

	logger.Info("Application done on command line.")
}
