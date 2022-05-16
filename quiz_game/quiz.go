package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "quiz_game/problems.csv", "a csv file in the format of 'questions,answer'")

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

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect")
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
