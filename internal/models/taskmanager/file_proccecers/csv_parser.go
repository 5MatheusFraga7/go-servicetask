package models

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type CsvParser struct {
}

func (cp CsvParser) GetDataFromInternalFile() [][]string {
	file, err := os.Open("data2.csv")
	// Checks for the error
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	defer file.Close()
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading records")
	}

	return records
}
