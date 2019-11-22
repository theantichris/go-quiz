// go-quiz is a simple quiz game written in Go.
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a CSV file in the form of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Printf("Could not open file %q: %v\n", *csvFilename, err)
		os.Exit(1)
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Could not read file %q: %v\n", *csvFilename, err)
		os.Exit(1)
	}

	for _, record := range records {
		fmt.Printf("%s\n", record[0])
	}
}
