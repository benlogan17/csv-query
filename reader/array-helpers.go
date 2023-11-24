package reader

func IndexOf(arr []string, value string) int {
	for index, item := range arr {
		if item == value {
			return index
		}
	}
	return -1
}
