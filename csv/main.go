package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("problems.csv")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	quizS, _ := reader.ReadAll()

	correct := 0

	for _, e := range quizS {
		quiz, answer := e[0], e[1]

		fmt.Printf("%v?", quiz)

		var userAnswer string
		fmt.Scan(&userAnswer)

		if userAnswer == answer {
			correct++
		}
	}
	fmt.Printf("Result: %v/%v\n", correct, len(quizS))
}
