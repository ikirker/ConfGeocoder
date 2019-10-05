package main

import (
	"fmt"
	"github.com/codingsince1985/geo-golang/openstreetmap"
)

func AddGeoCoords(c *Conference) {
	geocoder := openstreetmap.Geocoder()
	if (c.Location != "") &&
		(c.GeoCoords.Latitude == 0) &&
		(c.GeoCoords.Longitude == 0) &&
		(c.GeoCoords.Altitude == 0) {

		cityLocation, err := geocoder.Geocode(c.Location)
		fmt.Printf("%+v\n", cityLocation)
		if err != nil {
			panic(err)
		}
		c.GeoCoords.Latitude = cityLocation.Lat
		c.GeoCoords.Longitude = cityLocation.Lng
	}
}

func AddSeriesGeoCoords(cs *ConferenceSeries) {
	if cs.Next != nil {
		AddGeoCoords(cs.Next)
	}
	if cs.Last != nil {
		AddGeoCoords(cs.Last)
	}
}
