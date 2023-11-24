package commands

import (
	"demo/concurrency/functions"
	"demo/concurrency/reader"
	"demo/concurrency/types"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetGroupByColumn(input string, location *int) string {
	groupIndex := strings.Index(strings.ToLower(input), "group by") + 8
	columnToGroupBy := reader.GetNextWord(input, []string{" "}, &groupIndex)
	*location = groupIndex
	return columnToGroupBy
}

func getUniqueGroupByValues(data *types.Data, columnToGroupBy string, unqiueGroupByValues *[]string) {
	for _, row := range *data.Rows {
		if !reader.Contains(*unqiueGroupByValues, (*row)[columnToGroupBy]) {
			*unqiueGroupByValues = append(*unqiueGroupByValues, (*row)[columnToGroupBy])
		}
	}
}

func addGroupByColumn(data *types.Data, columnToGroupBy string, rows *[]*map[string]string) {
	unqiueGroupByValues := []string{}

	getUniqueGroupByValues(data, columnToGroupBy, &unqiueGroupByValues)

	for _, value := range unqiueGroupByValues {
		row := map[string]string{}

		row[columnToGroupBy] = value
		*rows = append(*rows, &row)
	}
}

func performFunction(data *types.Data, columnToGroupBy string, aggColumn string) *map[string]int {
	index := 0
	aggFunction := reader.GetNextWord(aggColumn, []string{"("}, &index)
	sumColumn := reader.GetNextWord(aggColumn, []string{")"}, &index)

	switch strings.ToLower(aggFunction) {
	case "count":
		countDataChan := make(chan *map[string]int)
		go functions.Count(data, columnToGroupBy, countDataChan)
		return <-countDataChan
	case "avg":
		return functions.Avg(data, columnToGroupBy, sumColumn)
	case "sum":
		sumDataChan := make(chan *map[string]int)
		go functions.Sum(data, columnToGroupBy, sumColumn, sumDataChan)
		return <-sumDataChan
	default:
		fmt.Printf("Unrecognised agg function: %s\n", aggFunction)
		os.Exit(1)
	}

	return nil
}

func SelectWithGroupBy(data *types.Data, columnToGroupBy string, columns *[]string, aggIndexes *[]int) *types.Data {
	rows := []*map[string]string{}
	aggData := types.Data{
		Header: columns,
		Rows:   &rows,
	}

	addGroupByColumn(data, columnToGroupBy, &rows)

	for _, aggIndex := range *aggIndexes {
		aggColumn := (*columns)[aggIndex]
		aggValues := performFunction(data, columnToGroupBy, aggColumn)

		for _, row := range *aggData.Rows {
			key := (*row)[columnToGroupBy]
			value := (*aggValues)[key]
			(*row)[aggColumn] = strconv.Itoa(value)
		}

	}
	return &aggData
}
