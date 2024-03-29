package main

import (
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, 301)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	cs := make([]UrlMap, 5)

	err := yaml.Unmarshal(yml, &cs)
	if err != nil {
		return nil, err
	}
	spew.Dump(cs)

	urlsMap := getMap(cs)

	return MapHandler(urlsMap, fallback), nil
}

func TOMLHandler(tml string, fallback http.Handler) (http.HandlerFunc, error) {
	cs := make(map[string][]UrlMap)

	_, err := toml.Decode(tml, &cs)
	if err != nil {
		return nil, err
	}
	spew.Dump(cs)
	urlsMap := getMap(cs["urls"])

	return MapHandler(urlsMap, fallback), nil
}

func getMap(cs []UrlMap) map[string]string {
	urlsMap := make(map[string]string)
	for _, c := range cs {
		urlsMap[c.Path] = c.URL
	}

	return urlsMap
}

type UrlMap struct {
	URL  string `yaml:"url" toml:"url"`
	Path string `yaml:"path" toml:"path"`
}
