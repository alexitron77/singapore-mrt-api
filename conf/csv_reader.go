package conf

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func ReadCSV(source string) [][]string {
	file, _ := filepath.Abs(source)
	csvFile, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully open file")
		defer csvFile.Close()
	}

	records, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	return records
}
