package commands

import (
	"demo/concurrency/reader"
	"demo/concurrency/types"
)

func createRows(indexRange []int, producer chan *map[string]string, oldRows *[]*map[string]string, columns []string) {
	for index := indexRange[0]; index < indexRange[1]; index++ {
		myMap := map[string]string{}
		for _, column := range columns {
			myMap[column] = (*(*oldRows)[index])[column]
		}
		producer <- &myMap
	}
}

func asyncSelectColumns(newData *types.Data, oldRows *[]*map[string]string) {
	producer := make(chan *map[string]string)
	dataLength := len(*newData.Rows)
	indexRange1 := []int{0, dataLength / 2}
	indexRange2 := []int{dataLength / 2, dataLength}
	go createRows(indexRange1, producer, oldRows, *newData.Header)
	go createRows(indexRange2, producer, oldRows, *newData.Header)
	current := 0
	for current < dataLength {
		value, ok := <-producer
		if ok == false {
			break
		}
		(*newData.Rows)[current] = value
		current++
	}
	close(producer)
}

func selectData(data *types.Data, columns []string, indexes []int) *types.Data {
	newRows := make([]*map[string]string, len(*data.Rows))
	newData := types.Data{
		Header: &columns,
		Rows:   &newRows,
	}

	asyncSelectColumns(&newData, data.Rows)

	return &newData
}

func processRenames(data *types.Data, rename *map[string]string, columns *[]string) {
	for key, value := range *rename {
		for _, row := range *data.Rows {
			(*row)[value] = (*row)[key]
		}

		if index := reader.IndexOf(*columns, key); index > -1 {
			(*columns)[index] = value
		}
	}
}

func Select(data *types.Data, input string, location *int) *types.Data {
	rename := map[string]string{}
	columns := getColumns(input, &rename)

	aggIndexes := reader.DoesAnyContain(columns, '(')

	if len(aggIndexes) > 0 {
		columnToGroupBy := GetGroupByColumn(input, location)
		aggData := SelectWithGroupBy(data, columnToGroupBy, &columns, &aggIndexes)
		processRenames(aggData, &rename, &columns)
		return aggData
	}

	var indexes []int
	if columns[0] == "*" {
		indexes = getIndexesOfHeader(*data.Header)
		columns = *data.Header
	} else {
		indexes = getIndexesOfColumns(*data.Header, columns)
	}

	data = selectData(data, columns, indexes)

	processRenames(data, &rename, &columns)

	return data
}
