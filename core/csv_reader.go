package core

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func ReadCSV(source string) [][]string {
	file := source
	csvFile, err := os.Open(file)

	if err != nil {
		fmt.Print(err)
		log.Fatal("Error while opening the file")
	} else {
		log.Print("Successfully open file")
		defer csvFile.Close()
	}

	records, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		log.Fatal("Error while reading the file")
	}

	return records
}
