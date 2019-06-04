package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/pflag"
)

var (
	mapFileYAML = pflag.StringP("yaml-map-file", "y", "", "urls mapping yaml file")
	mapFileTOML = pflag.StringP("toml-map-file", "t", "", "urls mapping toml file")
)

func main() {

	pflag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultHandler)

	var handler http.HandlerFunc

	if *mapFileYAML != "" {
		file, err := ioutil.ReadFile(*mapFileYAML)
		if err != nil {
			log.Fatalln("error read yaml file:", err)
		}
		handler, err = YAMLHandler(file, mux)
		if err != nil {
			log.Fatalln("cannot instantiate handler: ", err)
		}
	} else if *mapFileTOML != "" {
		file, err := ioutil.ReadFile(*mapFileTOML)
		if err != nil {
			log.Fatalln("error read yaml file:", err)
		}
		handler, err = TOMLHandler(string(file), mux)
		if err != nil {
			log.Fatalln("cannot instantiate handler: ", err)
		}
	} else {
		log.Fatalln("one of --yaml-map-file or toml-map-file is mandatory")
	}

	log.Fatal(http.ListenAndServe(":8080", handler))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Wolrd!")
}
