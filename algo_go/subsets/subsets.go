package subsets

func Subsets(nums []int) [][]int {
	subsets := make([][]int, 0)
	subsets = append(subsets, []int{})

	for _, num := range nums {
		for _, subset := range subsets {
			subsets = append(subsets, append(subset, num))
		}
	}
	return subsets
}
