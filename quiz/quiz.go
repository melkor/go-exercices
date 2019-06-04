package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/pflag"
)

var (
	CSVFileName = pflag.StringP("csv-file-name", "f", "", "input csv file")
	duration    = pflag.IntP("duration", "d", 30, "total quizz duration")
)

func main() {

	pflag.Parse()

	if *CSVFileName == "" {
		fmt.Println("csv-file-name is empty")
		os.Exit(1)
	}

	f, err := os.Open(*CSVFileName)

	if err != nil {
		fmt.Println("error open '", CSVFileName, " : ", err)
		os.Exit(1)
	}

	r := csv.NewReader(f)
	scanner := bufio.NewScanner(os.Stdin)

	goodResponses := 0
	badResponses := 0

	totalDuration := time.Duration(*duration) * time.Second
	timer := time.NewTimer(totalDuration)

	responseChannel := make(chan string)

mainloop:
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		go func() {
			fmt.Println("what ", record[0], ", Sir?")
			scanner.Scan()
			if err = scanner.Err(); err != nil {
				log.Fatal(err)
			}

			responseChannel <- strings.TrimSpace(scanner.Text())
		}()

		select {
		case <-timer.C:
			fmt.Println("Quizz is over")
			break mainloop

		case response := <-responseChannel:
			if response == record[1] {
				goodResponses++
			} else {
				badResponses++
			}
		}
	}

	fmt.Println("good responses :", goodResponses)
	fmt.Println("bad responses :", badResponses)

}
