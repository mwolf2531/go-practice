package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "quiz_game/problems.csv", "a csv file in the format of 'questions,answer'")

	timeLimit := flag.Int("limit", 30, "the time limit for the quiz")
	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		fmt.Println("Cannot read file")
	}

	csvReader := csv.NewReader(file)
	lines, err := csvReader.ReadAll()

	if err != nil {
		fmt.Println("Cannot read file")
	}

	problems := parseLines(lines)

	correct := 0

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.a {
				correct++
				fmt.Println("Correct!")
			} else {
				fmt.Println("Incorrect")
			}
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}
