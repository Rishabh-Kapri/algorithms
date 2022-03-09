package two_pointers

// Time: O(n), Space: O(1)
func TwoSum(input []int, expected int) (int, int) {

    leftPtr := 0
    rightPtr := len(input) - 1

    for leftPtr < rightPtr {
        sum := input[leftPtr] + input[rightPtr]
        if (sum == expected) {
            return input[leftPtr], input[rightPtr]
        } else if sum > expected {
            rightPtr--
        } else if sum < expected {
            leftPtr++
        }
    }
    return -1, -1
}
