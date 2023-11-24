package reader

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCSV(filename string) [][]string {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Could not open file")
		os.Exit(1)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()

	if err != nil {
		fmt.Println("Could not read file")
		os.Exit(1)
	}

	return data
}
