package main

import (
	"github.com/ty-edelweiss/csv2geojson"
	"github.com/ty-edelweiss/csv2geojson/log"
)

func main() {
	var opts csv2geojson.Options
	args := opts.Parse()

	log.Init(opts.Verbose)

	converter := csv2geojson.NewConverter(args[0], &opts)

	converter.Do()
}
