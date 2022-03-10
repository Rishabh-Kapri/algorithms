package sliding_window

// Maximum average contiguous subarray
/**
* keep on increasing the end index until the specified window size is reached
* keep on increasing the end while also increasing the start to maintain window size
 */
func MaxAvg(input []int, k int) float64 {
	start := 0
	sum := 0

	average := 0.0
	for i, num := range input {
		sum += num
		if i >= k-1 {
			newAvg := float64(sum) / float64(k)
			if newAvg > average {
				average = newAvg
			}
			sum -= input[start]
			start += 1
		}
	}
	return average
}
