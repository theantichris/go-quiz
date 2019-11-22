// go-quiz is a simple quiz game written in Go.
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

const csvFilenameFlag = "csv"
const csvFilenameDefault = "problems.csv"
const csvFilenameHelper = "a CSV file in the form of 'question,answer'"

type problem struct {
	question string
	answer   string
}

func main() {
	csvFilename := flag.String(csvFilenameFlag, csvFilenameDefault, csvFilenameHelper)
	timelimit := flag.Int("timer", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		handleError(fmt.Sprintf("Could not open file %q: %v\n", *csvFilename, err))
	}

	fileReader := csv.NewReader(file)
	lines, err := fileReader.ReadAll()
	if err != nil {
		handleError(fmt.Sprintf("Could not read file %q: %v\n", *csvFilename, err))
	}

	numberCorrect := 0
	problems := makeProblems(lines)
	timer := time.NewTimer(time.Duration(*timelimit) * time.Second)

questionLoop:
	for _, problem := range problems {
		fmt.Printf("%s = ", problem.question)

		answerChan := make(chan string)
		go func() {
			var answer string
			_, err := fmt.Scanf("%s\n", &answer)
			if err != nil {
				handleError(fmt.Sprintf("Could not read input: %v", err))
			}
			answerChan <- answer
		}()

		select {
		case answer := <-answerChan:
			if answer == problem.answer {
				numberCorrect++
			}
		case <-timer.C:
			fmt.Println()
			break questionLoop
		}
	}

	fmt.Printf("Your correct answers: %d of %d\n", numberCorrect, len(lines))
}

func makeProblems(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		problems[i] = problem{line[0], line[1]}
	}

	return problems
}

func handleError(message string) {
	fmt.Println(message)
	os.Exit(1)
}
