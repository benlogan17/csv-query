package reader

func Contains(arr []string, value string) bool {
	for _, item := range arr {
		if item == value {
			return true
		}
	}
	return false
}

func isSpace(current byte) bool {
	return current == ' '
}

func DoesAnyContain(arr []string, char byte) []int {
	indexes := []int{}
	for itemIndex, item := range arr {
		for _, charItem := range item {
			if char == byte(charItem) {
				indexes = append(indexes, itemIndex)
			}
		}
	}
	return indexes
}

func GetNextWord(input string, terminators []string, index *int) string {
	word := ""
	for *index < len(input) {
		if isSpace(input[*index]) && len(word) == 0 {
			*index++
		} else if Contains(terminators, string(input[*index])) {
			*index++
			return word
		} else {
			word += string(input[*index])
			*index++
		}
	}
	return word
}
