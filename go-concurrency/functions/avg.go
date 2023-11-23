package functions

import (
	"demo/concurrency/types"
)

func Avg(data *types.Data, columnToGroupBy string, sumColumn string) *map[string]int {
	countDataChan := make(chan *map[string]int)
	sumDataChan := make(chan *map[string]int)

	go Count(data, columnToGroupBy, countDataChan)
	go Sum(data, columnToGroupBy, sumColumn, sumDataChan)

	countData := <-countDataChan
	sumData := <-sumDataChan

	avgValues := map[string]int{}

	for key, value := range *countData {
		avgValues[key] = (*sumData)[key] / value
	}
	return &avgValues
}
