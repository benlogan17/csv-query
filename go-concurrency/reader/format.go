package reader

import "demo/concurrency/types"

func ConverToDataMap(data [][]string) types.Data {
	headers := data[0]

	var rows []*map[string]string

	rows = make([]*map[string]string, len(data)-1)

	for rowIndex := 1; rowIndex < len(data); rowIndex++ {
		addRow := map[string]string{}
		for itemIndex, header := range headers {
			addRow[header] = data[rowIndex][itemIndex]
		}
		rows[rowIndex-1] = &addRow
	}

	return types.Data{
		Header: &headers,
		Rows:   &rows,
	}
}
