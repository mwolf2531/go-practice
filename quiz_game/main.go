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

	if err != nil {
		fmt.Println("Cannot read file")
	}

	//correct := 0

	for {
		record, err := csvReader.Read()
		if err != nil {
			break
		}
		fmt.Printf("Problem #%s Answer %s\n", record[0], record[1])
	}

}
