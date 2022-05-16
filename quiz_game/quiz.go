package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("quiz_game/problems.csv")

	if err != nil {
		fmt.Println("Cannot read file")
	}

	csvReader := csv.NewReader(file)
	lines, err := csvReader.ReadAll()

	if err != nil {
		fmt.Println("Cannot read file")
	}

	problems := parseLines(lines)
	fmt.Println(problems)
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

// for {
// 	record, err := csvReader.Read()
// 	if err != nil {
// 		break
// 	}
// 	fmt.Printf("Problem #%s Answer %s\n", record[0], record[1])
// }
