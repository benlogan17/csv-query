package commands

import (
	"strings"

	"demo/concurrency/reader"
)

func generateColumnsArray(input string) []string {
	indexOfFROM := strings.Index(strings.ToLower(input), "from")
	length := strings.Count(input[0:indexOfFROM], ",") + 1
	return make([]string, length)
}

func getColumns(input string, rename *map[string]string) []string {
	index := strings.Index(input, " ")

	columns := generateColumnsArray(input)
	columnsindex := 0
	terminators := []string{",", " "}
	renameFlag := false
	finished := true

	for finished {
		word := reader.GetNextWord(input, terminators, &index)
		if word == "as" {
			renameFlag = true
		} else if renameFlag {
			original := columns[columnsindex-1]
			(*rename)[original] = word
			renameFlag = false
		} else if columnsindex < len(columns) {
			columns[columnsindex] = word
			columnsindex++
		} else {
			finished = false
		}
	}
	return columns
}
