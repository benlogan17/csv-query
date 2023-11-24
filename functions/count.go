package functions

import (
	"demo/concurrency/types"
)

func Count(data *types.Data, columnToGroupBy string, countChan chan *map[string]int) {
	uniqueCount := map[string]int{}

	for _, row := range *data.Rows {
		value := (*row)[columnToGroupBy]
		if _, ok := uniqueCount[value]; ok {
			uniqueCount[value] += 1
		} else {
			uniqueCount[value] = 1
		}
	}

	countChan <- &uniqueCount
}
