package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	// flag.Parse()

	file, err := os.Open("quiz_game/problems.csv")

	if err != nil {
		fmt.Println("Cannot read file 1")
	}

	csvReader := csv.NewReader(file)

	//lines, err := csvReader.ReadAll()

	if err != nil {
		fmt.Println("Cannot read file 2")
	}

	for {
		record, err := csvReader.Read()
		if err != nil {
			break
		}
		fmt.Printf("Question: %s Answer %s\n", record[0], record[1])
	}

}
