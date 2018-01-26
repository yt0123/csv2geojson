package geo

import (
	"crypto/sha1"
	"errors"
	"io"
	"strconv"

	"github.com/paulmach/go.geojson"
	"github.com/ty-edelweiss/csv2geojson/log"
)

type Point []float64
type LineString [][]float64
type Polygon [][][]float64

func BuildPointCollection(longitude string, latitude string, columns []int, headers []string, records [][]string) ([]byte, error) {
	fc := geojson.NewFeatureCollection()

	for _, record := range records {
		coord, err := ParseCoordinate(columns, record)
		if err != nil {
			log.Logger.Debug(err)
			continue
		}

		feature := geojson.NewPointFeature(coord)

		properties := ParseProperties(headers, record, longitude, latitude)
		for key, property := range properties {
			feature.SetProperty(key, property)
		}

		fc.AddFeature(feature)
	}

	return fc.MarshalJSON()
}

func BuildLineStringCollection(longitude string, latitude string, index int, columns []int, headers []string, records [][]string) ([]byte, error) {
	fc := geojson.NewFeatureCollection()

	tmp := make(map[string]LineString)
	tmps := make(map[string]PropertyCollections)
	for _, record := range records {
		coord, err := ParseCoordinate(columns, record)
		if err != nil {
			log.Logger.Debug(err)
			continue
		}

		key := record[index]
		log.Logger.WithField("key", key).Debug("Record key value is following.")

		properties := ParseProperties(headers, record, longitude, latitude)

		tmp[key] = append(tmp[key], coord)
		log.Logger.WithField("contents", tmp).Debug("Key data contents is following.")

		if _, ok := tmps[key]; !ok {
			tmps[key] = PropertyCollections{}
		}
		tmps[key].AppendProperties(properties)
		log.Logger.WithField("properties", tmps).Debug("Key data properties is following.")
	}

	for id, coords := range tmp {
		feature := geojson.NewLineStringFeature(coords)

		feature.SetProperty("hash_", ParseHash(id))

		for key, pc := range tmps[id] {
			feature.SetProperty(key, pc)
		}

		fc.AddFeature(feature)
	}

	return fc.MarshalJSON()
}

func BuildPolygonCollection(longitude string, latitude string, index int, columns []int, headers []string, records [][]string) ([]byte, error) {
	fc := geojson.NewFeatureCollection()

	tmp := make(map[string]LineString)
	tmps := make(map[string]PropertyCollections)
	for _, record := range records {
		coord, err := ParseCoordinate(columns, record)
		if err != nil {
			log.Logger.Warn(err)
			continue
		}

		key := record[index]

		properties := ParseProperties(headers, record, longitude, latitude)

		tmp[key] = append(tmp[key], coord)
		log.Logger.WithField("contents", tmp).Debug("Key data contents is following.")

		if _, ok := tmps[key]; !ok {
			tmps[key] = PropertyCollections{}
		}
		tmps[key].AppendProperties(properties)
		log.Logger.WithField("properties", tmps).Debug("Key data properties is following.")
	}

	for id, coords := range tmp {

		polygon, err := ParsePolygon(coords)
		if err != nil {
			log.Logger.Warn(err)
			continue
		}

		feature := geojson.NewPolygonFeature(polygon)

		feature.SetProperty("hash_", ParseHash(id))

		for key, prop := range tmps[id] {
			feature.SetProperty(key, prop)
		}

		fc.AddFeature(feature)
	}

	return fc.MarshalJSON()
}

func ParseCoordinate(columns []int, record []string) (Point, error) {
	if len(columns) == 1 {
		return []float64{}, errors.New("Coordinate format is invalid")
	}

	lon, err := strconv.ParseFloat(record[columns[0]], 64)
	if err != nil {
		return []float64{}, err
	}
	lat, err := strconv.ParseFloat(record[columns[1]], 64)
	if err != nil {
		return []float64{}, err
	}

	return Point{lon, lat}, nil
}

func ParsePolygon(lines ...LineString) (Polygon, error) {
	if len(lines) > 2 {
		return Polygon{}, errors.New("Polygon parse arguments is too many")
	}

	polygon := Polygon{}
	for _, line := range lines {
		if len(line) < 3 {
			return Polygon{}, errors.New("Coordinates format is invalid for polygon")
		}

		if line[len(line)-1][0] == line[0][0] && line[len(line)-1][1] == line[0][1] {
			polygon = append(polygon, line)
		} else {
			closed := append(line, line[0])
			polygon = append(polygon, closed)
		}
	}

	return polygon, nil
}

func ParseHash(key string) []byte {
	hash := sha1.New()

	io.WriteString(hash, key)

	buf := hash.Sum(nil)
	log.Logger.WithField("hash", buf).Info("Convert key to hash buffer done")

	return buf
}
