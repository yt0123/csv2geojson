package geo

import (
	"errors"

	"github.com/ty-edelweiss/csv2geojson/log"
)

func Build(geotype string, longitude string, latitude string, key string, headers []string, records [][]string) ([]byte, error) {
	columns, err := collect(longitude, latitude, headers)
	if err != nil {
		return []byte{}, err
	}

	log.Logger.WithField("geotype", geotype).Info("Selected geometry type is following.")
	switch geotype {
	case "Point":
		return BuildPointCollection(longitude, latitude, columns, headers, records)
	case "LineString":
		index, err := check(key, headers)
		if err != nil {
			return []byte{}, err
		}
		return BuildLineStringCollection(longitude, latitude, index, columns, headers, records)
	case "Polygon":
		index, err := check(key, headers)
		if err != nil {
			return []byte{}, err
		}
		return BuildPolygonCollection(longitude, latitude, index, columns, headers, records)
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
