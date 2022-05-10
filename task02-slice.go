package homework

func reverse(input []int64) (result []int64) {
	result = make([]int64, len(input), len(input))

	for i, r := 0, len(result)-1; i < len(input); i, r = i+1, r-1 {
		result[r] = input[i]
	}

	return
}
