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
)

func main() {

	pflag.Parse()

	if *mapFileYAML == "" {
		log.Fatalln("--yaml-map-file is mandatory")
	}

	yamlFile, err := ioutil.ReadFile(*mapFileYAML)
	if err != nil {
		log.Fatalln("error read yaml file:", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultHandler)

	handler, err := YAMLHandler(yamlFile, mux)
	if err != nil {
		log.Fatalln("cannot instantiate handler: ", err)
	}
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Wolrd!")
}
