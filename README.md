# Conference GeoCoder

This is a proof-of-concept-level tool to convert the YAML files in @jamespjh's [collection of eResearch conference information] to GeoJSON that GitHub will render into a nice mapbox display.

It's written in Go and uses the modules mechanism, so if you have a Go compiler you can clone it and build it in-place.

```
git clone https://github.com/ikirker/ConfGeocoder.git
cd ConfGeocoder
go build .
./ConfGeocoder some/dir/*.yml >collection.geojson
```

[Here's an example I made earlier.](https://gist.github.com/ikirker/a909b4d4205aca3188a55ef1e6a54d70)


