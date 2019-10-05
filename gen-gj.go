package main

import (
	"encoding/json"
	"fmt"
	"github.com/codingsince1985/geo-golang/openstreetmap"
	"os"
)

type GJFeatureCollection struct {
	Type     string      `json:"type"` // "FeatureCollection"
	Features []GJFeature `json:"features"`
}

// Note: this is only as flexible as it needs to be to make points
// If you were going to make a proper thing that could handle regions and stuff,
//  you'd need more types, maybe with an interface.
type GJFeature struct {
	Type       string            `json:"type"` // "Feature"
	Geometry   GJPoint           `json:"geometry"`
	Properties map[string]string `json:"properties"`
}

type GJPoint struct {
	Type        string    `json:"type"` // "Point"
	Coordinates []float64 `json:"coordinates"`
}

func GenerateGeoJSON(conferenceSerieses []ConferenceSeries) []byte {

	f := GJFeatureCollection{Type: "FeatureCollection"}
	f.Features = []GJFeature{}

	for _, v := range conferenceSerieses {
		// We're only interested in conferences with information we can plot
		// We could add on the previous or last conferences but that logic is finicky and I'm tired
		if v.Next == nil {
			continue
		}

		f.Features = append(f.Features, MakeFeatureForConference(v, *v.Next))
	}

	jsonBytes, err := json.Marshal(f)
	if err != nil {
		panic(err)
	}
	return jsonBytes
}

// It's just easier
type Marker map[string]string

// The basic styling
func NewMarker() Marker {
	var m map[string]string
	m = make(map[string]string)
	m["marker-color"] = "#ffccdd"
	m["marker-size"] = "small"
	return m
}

func aOb(a string, b string) string {
	if a != "" {
		return a
	} else {
		return b
	}
}

func MakeFeatureForConference(series ConferenceSeries, instance Conference) GJFeature {
	// NB: all non-special marker fields and values get rendered as HTML
	marker := NewMarker()
	feature := GJFeature{}

	if instance.Date != "" {
		marker["Date"] = instance.Date
	} else if instance.DateFrom != "" {
		marker["Date From"] = instance.DateFrom
		if instance.DateTo != "" {
			marker["Date To"] = instance.DateTo
		}
	}

	marker["Name"] = aOb(instance.Name, aOb(series.LongName, series.Name))
	marker["Frequency"] = series.Frequency
	marker["Link"] = "<a href=\"" + aOb(instance.Link, series.Link) + "\">" + aOb(instance.Link, series.Link) + "</a>"
	marker["Location"] = aOb(instance.Location, series.Location)

	mapPoint := GJPoint{
		Type: "Point",
		Coordinates: []float64{
			// Notice these are flipped from the order
			//  they're usually in, for some reason.
			instance.GeoCoords.Longitude,
			instance.GeoCoords.Latitude,
		},
	}

	feature.Type = "Feature"
	feature.Properties = marker
	feature.Geometry = mapPoint
	return feature
}

func main() {
	geocoder := openstreetmap.Geocoder()

	confSeriesFiles := os.Args[1:]
	confSerieses := []ConferenceSeries{}
	for _, v := range confSeriesFiles {
		fmt.Fprintf(os.Stderr, "Reading %s...\n", v)
		confSeries := ReadConferenceSeries(v)
		if confSeries.Next != nil {
			if confSeries.Next.Location != "" {
				cityLocation, err := geocoder.Geocode(confSeries.Next.Location)
				if err != nil {
					panic(err)
				}
				confSeries.Next.GeoCoords.Latitude = cityLocation.Lat
				confSeries.Next.GeoCoords.Longitude = cityLocation.Lng
			}
		}
		confSerieses = append(confSerieses, confSeries)
	}
	output := string(GenerateGeoJSON(confSerieses))
	fmt.Printf("%s\n", output)
}
