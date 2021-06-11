package core

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
)

func ReadCSV(source string) [][]string {
	file, _ := filepath.Abs(source)
	csvFile, err := os.Open(file)

	if err != nil {
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
