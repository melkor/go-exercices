package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultHandler)
	pathsToUrls := map[string]string{
		"/lol":  "https://godoc.org/github.com/gophercises/urlshort",
		"/bite": "https://godoc.org/gopkg.in/yaml.v2",
	}
	//MapHandler(pathsToUrls, mux)
	log.Fatal(http.ListenAndServe(":8080", MapHandler(pathsToUrls, mux)))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Wolrd!")
}
