// go-quiz is a simple quiz game written in Go.
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

const csvFilenameFlag = "csv"
const csvFilenameDefault = "problems.csv"
const csvFilenameHelper = "a CSV file in the form of 'question,answer'"

func main() {
	csvFilename := flag.String(csvFilenameFlag, csvFilenameDefault, csvFilenameHelper)
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Printf("Could not open file %q: %v\n", *csvFilename, err)
		os.Exit(1)
	}

	fileReader := csv.NewReader(file)
	records, err := fileReader.ReadAll()
	if err != nil {
		fmt.Printf("Could not read file %q: %v\n", *csvFilename, err)
		os.Exit(1)
	}

	numberCorrect := 0
	for _, record := range records {
		fmt.Printf("%s: ", record[0])

		var userAnswer string
		_, err := fmt.Scanf("%s\n", &userAnswer)
		if err != nil {
			fmt.Printf("Could not read input: %v", err)
			os.Exit(0)
		}

		if userAnswer == record[1] {
			numberCorrect++
		}
	}

	fmt.Printf("Your correct answers: %d of %d\n", numberCorrect, len(records))
	os.Exit(0)
}
