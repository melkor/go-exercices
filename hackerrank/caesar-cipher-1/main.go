package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/pflag"
)

var (
	stringToCipher = pflag.StringP("string-to-cipher", "s", "", "string to cipher")
	numberRotation = pflag.IntP("number-rotations", "k", 0, "number of rotation")
)

func main() {
	pflag.Parse()

	if *stringToCipher == "" {
		panic("string-to-cipher is empty")
	}

	ciphered, err := cipher(*stringToCipher, *numberRotation)
	if err != nil {
		panic(err)
	}
	fmt.Println(ciphered)
}

func cipher(s string, k int) (string, error) {
	r := make([]string, 0, len(s))

	if k < 0 || k > 100 {
		return "", errors.New("k is invalid")
	}

	ss := strings.Split(s, "")

	if len(ss) < 1 || len(ss) > 100 {
		return "", errors.New("s is invalid")
	}

	for _, c := range ss {

		if c == "-" || c == "'" {
			r = append(r, c)
		} else {
			asciiPos := int(c[0])
			newAsciiPos := 0
			if asciiPos >= int('a') || asciiPos <= int('z') {
				newAsciiPos = asciiPos + k
				if newAsciiPos > int('z') {
					r = append(r, string(int('a')+newAsciiPos-int('z')-1))
				} else {
					r = append(r, string(newAsciiPos))
				}
			} else if asciiPos >= int('A') || asciiPos <= int('Z') {
				newAsciiPos = asciiPos + k
				if newAsciiPos > int('Z') {
					r = append(r, string(int('A')+newAsciiPos-int('Z')-1))
				} else {
					r = append(r, string(newAsciiPos))
				}
			} else {
				return "", errors.New("s is an invalid ascii string")
			}

		}
	}

	return strings.Join(r[:], ""), nil
}
