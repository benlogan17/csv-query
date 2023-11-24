package commands

import (
	"demo/concurrency/reader"
	"demo/concurrency/types"
	"fmt"
)

func Remove(index int, data *types.Data) {
	rows := *data.Rows
	rows[index] = rows[len(rows)-1]
	newRows := rows[:len(rows)-1]
	data.Rows = &newRows
}

func checkValueByOperation(operation string, value1 string, value2 string) bool {
	switch operation {
	case "==":
		return value1 != value2
	case ">":
		return value1 <= value2
	case "<":
		return value1 >= value2
	case "<=":
		return value1 > value2
	case ">=":
		return value1 < value2
	case "!=":
		return value1 == value2
	default:
		return false
	}
}

func Where(data *types.Data, input string, location *int) {
	spaceTerminator := []string{" "}
	field := reader.GetNextWord(input, spaceTerminator, location)
	operation := reader.GetNextWord(input, spaceTerminator, location)
	value := reader.GetNextWord(input, spaceTerminator, location)

	if !reader.Contains(*data.Header, field) {
		fmt.Println("Column not present in file")
	}

	for rowIndex := 0; rowIndex < len((*data.Rows)); {
		rowValue := (*(*data.Rows)[rowIndex])[field]
		if checkValueByOperation(operation, rowValue, value) {
			Remove(rowIndex, data)
		} else {
			rowIndex++
		}
	}
}
