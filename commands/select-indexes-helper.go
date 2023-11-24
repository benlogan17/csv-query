package commands

import "strings"

func getIndexesOfHeader(header []string) []int {
	indexes := make([]int, len(header))
	current := 0
	for i := 0; i < len(header); i++ {
		indexes[current] = i
		current++
	}
	return indexes
}

func isLowerEqual(a string, b string) bool {
	return strings.ToLower(a) == strings.ToLower(b)
}

func getIndexesOfColumns(headers []string, columns []string) []int {
	indexes := make([]int, len(columns))
	current := 0

	for _, column := range columns {
		for index, header := range headers {
			if isLowerEqual(column, header) {
				indexes[current] = index
				current++
			}
		}
	}
	return indexes
}
