package commands

import (
	"demo/concurrency/reader"
	"demo/concurrency/types"
	"sort"
	"strings"
)

func Order(data *types.Data, location *int, input string) {
	*location += 3
	spaceTerminator := []string{" "}
	field := reader.GetNextWord(input, spaceTerminator, location)
	direction := strings.ToLower(reader.GetNextWord(input, spaceTerminator, location))

	rows := data.Rows
	sort.Slice(*rows, func(i, j int) bool {
		if direction == "desc" {
			return (*(*rows)[i])[field] > (*(*rows)[j])[field]
		}
		return (*(*rows)[i])[field] < (*(*rows)[j])[field]
	})
}
