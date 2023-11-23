package output

import (
	"fmt"
	"strings"

	"demo/concurrency/types"
)

func getColumnLengths(data *types.Data) map[string]int {
	columnLengths := map[string]int{}

	for _, header := range *data.Header {
		columnLengths[header] = len(header)
	}

	for _, row := range *data.Rows {
		for key, value := range *row {
			if len(value) > columnLengths[key] {
				columnLengths[key] = len(value)
			}
		}
	}

	return columnLengths
}

func addColumnItem(rowString *string, item string, numberOfSpaces int) {
	if numberOfSpaces%2 != 0 {
		halfSpaces := (numberOfSpaces - 1) / 2
		*rowString += " " + strings.Repeat(" ", halfSpaces+1) +
			item + strings.Repeat(" ", halfSpaces) + " |"
	} else {
		halfSpaces := numberOfSpaces / 2
		*rowString += " " + strings.Repeat(" ", halfSpaces) +
			item + strings.Repeat(" ", halfSpaces) + " |"
	}
}

func displayRow(columns []string, row map[string]string, columnLengths map[string]int) {
	rowString := "|"
	var numberOfSpaces int

	for _, column := range columns {
		numberOfSpaces = columnLengths[column] - len(row[column])
		addColumnItem(&rowString, row[column], numberOfSpaces)
	}
	fmt.Println(rowString)
}

func displayHeader(row []string, columnLengths map[string]int) {
	rowString := "|"
	var numberOfSpaces int

	for _, value := range row {
		numberOfSpaces = columnLengths[value] - len(value)
		addColumnItem(&rowString, value, numberOfSpaces)
	}
	top := strings.Repeat("-", len(rowString))
	fmt.Println(top)
	fmt.Println(rowString)
	fmt.Println(top)
}

func DisplayData(data *types.Data) {
	columnLengths := getColumnLengths(data)

	displayHeader(*data.Header, columnLengths)

	for _, row := range *data.Rows {
		displayRow(*data.Header, *row, columnLengths)
	}
}
