package geo

import (
	"errors"

	"github.com/ty-edelweiss/csv2geojson/log"
)

var (
	report = log.NewReport()
	logger = log.AppLogger.SetName("geo")
)

func Build(geotype string, longitude string, latitude string, key string, headers []string, records [][]string, limit int) ([]byte, error) {
	columns, err := collect(longitude, latitude, headers)
	if err != nil {
		return []byte{}, err
	}

	logger.WithField("geotype", geotype).Info("Selected geometry type is following.")
	switch geotype {
	case "Point":
		report.NewProgressBar("Build GeoJSON (Point) - ", len(records))
		fc := BuildPointCollection(longitude, latitude, columns, headers, records, limit)
		fc.CRS = NamedCRS(4326)
		report.ProgressDone()
		return fc.MarshalJSON()
	case "LineString":
		index, err := check(key, headers)
		if err != nil {
			return []byte{}, err
		}
		report.NewProgressBar("Build GeoJSON (LineString) - ", len(records))
		fc := BuildLineStringCollection(longitude, latitude, index, columns, headers, records, limit)
		fc.CRS = NamedCRS(4326)
		report.ProgressDone()
		return fc.MarshalJSON()
	case "Polygon":
		index, err := check(key, headers)
		if err != nil {
			return []byte{}, err
		}
		report.NewProgressBar("Build GeoJSON (Polygon) - ", len(records))
		fc := BuildPolygonCollection(longitude, latitude, index, columns, headers, records, limit)
		fc.CRS = NamedCRS(4326)
		report.ProgressDone()
		return fc.MarshalJSON()
	default:
		return []byte{}, errors.New("Geometry type is invalid")
	}
}

func collect(longitude string, latitude string, headers []string) ([]int, error) {
	var columns []int
	for index, header := range headers {
		if header == longitude || header == latitude {
			columns = append(columns, index)
		}
	}

	if len(columns) == 0 {
		return columns, errors.New("Target coordinate column is not exist")
	}

	return columns, nil
}

func check(key string, headers []string) (int, error) {
	index := -1
	for i, header := range headers {
		if header == key {
			index = i
			break
		}
	}

	if index == -1 {
		return index, errors.New("Key column is not exist")
	}

	return index, nil
}
