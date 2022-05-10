package homework

func average(input [15]float32) (result float32) {
	for _, item := range input {
		result += item
	}

	result /= float32(len(input))

	return

	/*
		@DzmitryYafremenka  wrote idiot's style tests that lead to write ^^^^ that bullshit ^^^^
		i.e. float comparison with equal\not equal is definitely bad practice

		revLen := 1.0 / float32(len(input))

			for _, item := range input {
				result += item * revLen
			}

			return*/
}
