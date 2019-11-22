// go-quiz is a simple quiz game written in Go.
package main

import (
	"flag"
	"fmt"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a CSV file in the form of 'question,answer'")
	flag.Parse()

	fmt.Println(*csvFilename)
}
