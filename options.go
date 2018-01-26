package csv2geojson

import (
	flags "github.com/jessevdk/go-flags"
	"os"
)

type Options struct {
	Output    string `short:"o" long:"output" description:"Set output path for converted geojson file"`
	Type      string `short:"t" long:"type" default:"Point" description:"Set geometry type for geojson file"`
	Key       string `short:"k" long:"key" description:"Set key column to join records"`
	Delimiter string `short:"d" long:"delimiter" default:"," description:"Set csv delimiter for imported csv file"`
	Longitude string `long:"lon" default:"longitude" description:"Set geometry coordinates for geojson file"`
	Latitude  string `long:"lat" default:"latitude" description:"Set geometry coordinates for geojson file"`
	Quotes    bool   `long:"quotes" description:"Check csv double quotes for imported csv file"`
	Preformat bool   `short:"p" long:"preformat" description:"Output preformatted geojson file"`
	Verbose   bool   `short:"v" long:"verbose" description:"Show verbose debug information"`
}

func (opts *Options) Parse() []string {
	parser := flags.NewParser(opts, flags.Default)
	parser.Name = "csv2geojson"
	parser.Usage = "[OPTIONS] PATTERN [PATH]"

	args, _ := parser.Parse()
	if len(args) == 0 {
		if !parser.FindOptionByLongName("help").IsSet() {
			parser.WriteHelp(os.Stdout)
		}
		os.Exit(1)
	}

	return args
}
