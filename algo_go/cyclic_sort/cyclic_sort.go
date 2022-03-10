package cyclic_sort

func MissingNumber(input []int) int {
	// [9, 6, 4, 2, 3, 5, 7, 0, 1]
	// start from the first index
	index := 0
	for index < len(input) {
		num := input[index]
		if num == index || num >= len(input) {
			// if num is at correct index or num is greater than or equal to
			// len, continue
			index++
		} else {
			// swap numbers
			swapNum := input[num]
			input[num] = num
			input[index] = swapNum
		}
	}

	for i, num := range input {
		if i != num {
			return i
		}
	}
	return len(input)
}
