package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// Note: we don't expect to have to output these as JSON;
//  but the JSON names might be useful for e.g. debugging dumps
type ConferenceSeries struct {
	Name      string       `yaml:"name";json:"name"`
	LongName  string       `yaml:"long-name";json:"long_name"`
	Frequency string       `yaml:"frequency";json:"frequency"`
	Link      string       `yaml:"link";json:"link"`
	Location  string       `yaml:"location";json:"location"`
	Next      *Conference  `yaml:"next";json:"next"`
	Previous  []Conference `yaml:"previous";json:"previous"`
	Last      *Conference  `yaml:"last";json:"last"`
	Interests []string     `yaml:"interests";json:"interests"`
}

type Conference struct {
	Name      string      `yaml:"name";json:"name"` // This should override the series name in many uses but may be empty
	Location  string      `yaml:"location";json:"location"`
	Link      string      `yaml:"link";json:"link"`
	GeoCoords GeoLocation `yaml:"geolocation";json:"geolocation"`
	DateFrom  string      `yaml:"date-from";json:"date_from"` // Yes technically this should be a *time.Time but let's not go there right now
	DateTo    string      `yaml:"date-to";json:"date_to"`     // Ditto
	Date      string      `yaml:"date";json:"date"`           // Ditto
}

type GeoLocation struct {
	Latitude  float64 `yaml:"latitude";json:"latitude"`
	Longitude float64 `yaml:"longitude";json:"longitude"`
	Altitude  float64 `yaml:"altitude";json:"altitude"` //  If you really want to tell us how high your conference is, I won't stop you
}

func ReadConferenceSeries(filename string) ConferenceSeries {
	yamlFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer yamlFile.Close()

	yamlBytes, err := ioutil.ReadAll(yamlFile)
	if err != nil {
		panic(err)
	}

	return ParseConferenceSeries(yamlBytes)
}

func ParseConferenceSeries(yamlBytes []byte) ConferenceSeries {
	confSeries := ConferenceSeries{}
	err := yaml.Unmarshal(yamlBytes, &confSeries)
	if err != nil {
		panic(err)
	}
	return confSeries
}
