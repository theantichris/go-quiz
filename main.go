// go-quiz is a simple quiz game written in Go.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a CSV file in the form of 'question,answer'")
	flag.Parse()

	data, err := ioutil.ReadFile(*csvFilename)
	if err != nil {
		fmt.Printf("Could not read file %q: %v\n", *csvFilename, err)
		os.Exit(1)
	}

	fmt.Print(string(data))
}
