package main

import (
	"fmt"
	"regexp"

	"github.com/spf13/pflag"
)

var (
	camlCaseString = pflag.StringP("caml-string", "s", "", "sequence of words in CamelCase")
)

func main() {
	pflag.Parse()

	if *camlCaseString == "" {
		panic("ERROR: string is empty (--s/--caml-string parameter)")
	}

	fmt.Println("Number of wrods in ", *camlCaseString, ":", countWords(*camlCaseString))
}

func countWords(s string) int {
	a := regexp.MustCompile("[A-Z]")

	return len(a.Split(s, -1))
}
