package functions

import (
	"demo/concurrency/types"
	"fmt"
	"os"
	"strconv"
)

func Sum(data *types.Data, columnToGroupBy string, sumColumn string, sumChan chan *map[string]int) {
	uniqueSum := map[string]int{}

	for _, row := range *data.Rows {
		key := (*row)[columnToGroupBy]
		value, err := strconv.Atoi((*row)[sumColumn])

		if err != nil {
			fmt.Printf("Column %s is wrong type for summing. Value is %s\n", sumColumn, (*row)[sumColumn])
			os.Exit(1)
		}

		if _, ok := uniqueSum[key]; ok {
			uniqueSum[key] += value
		} else {
			uniqueSum[key] = value
		}
	}

	sumChan <- &uniqueSum
}
