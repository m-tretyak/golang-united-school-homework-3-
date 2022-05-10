package homework

import (
	"sort"
)

func sortMapValues(input map[int]string) (result []string) {
	indexes := make([]int, 0, len(input))
	for index := range input {
		indexes = append(indexes, index)
	}

	sort.Ints(indexes)

	result = make([]string, 0, len(indexes))
	for index := range indexes {
		result = append(result, input[index])
	}

	return
}
