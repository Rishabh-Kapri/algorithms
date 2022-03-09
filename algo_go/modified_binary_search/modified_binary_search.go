package binary_search

func BinarySearch(nums []int, target int) int {
    start := 0
    end := len(nums) - 1

    for start <= end {
        mid := (start + end) / 2
        num := nums[mid]
        if num == target {
            return mid
        } else if num < target {
            start = mid + 1
        } else {
            end = mid - 1
        } 
    }
    return -1
}
