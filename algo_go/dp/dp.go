package dp

import (
    "fmt"
)

var Memo map[int]int
var TotalUsage map[int]int
var Tab map[int]int

func init() {
    Memo = make(map[int]int)
    TotalUsage = make(map[int]int)
    Tab = make(map[int]int)

    Memo[0] = 0
    Memo[1] = 1
    Tab[0] = 0
    Tab[1] = 1

}

func FibonacciMemo(num int) int {
    _, exists := Memo[num]

    if exists {
        TotalUsage[num] += 1
        return Memo[num]
    } else {
        Memo[num] = FibonacciMemo(num - 1) + FibonacciMemo(num - 2)
        return Memo[num]
    }
}

func FibonacciTab(num int) int {
    for i := 2; i <= num; i++ {
        Tab[i] = Tab[i - 1] + Tab[i - 2]
    }
    fmt.Println(Tab)
    return Tab[num]
}

func canPartition(nums []int) (bool, int) {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    if sum % 2 != 0 {
        return false, -1
    }
    return true, sum / 2
}

func PartEqualSumMemo(nums []int) bool {
    // [1, 5, 13, 5]
    // [1, 5, 5] [11]
    canPart, sum := canPartition(nums)

    if !canPart {
        return false
    }

    memo := make([][]int, len(nums))
    for i := 0; i < len(nums); i++ {
        memo[i] = make([]int, sum + 1)
        for j := 0; j <= sum; j++ {
            memo[i][j] = -1
        }
    }
    if partRecursive(nums, sum, 0, &memo) == 1 {
        return true
    } else {
        return false
    }
}

func partRecursive(nums []int, sum int, currentIndex int, memo *[][]int) int {
    // base check
    if sum == 0 {
        return 1
    }
    // base check
    if currentIndex >= len(nums) || len(nums) == 0 {
        return 0
    }

    if (*memo)[currentIndex][sum] == -1 {
        if nums[currentIndex] <= sum {
            // if sum is still not zero or index is still inbound increment index
            if partRecursive(nums, sum - nums[currentIndex], currentIndex + 1, memo) == 1 {
                (*memo)[currentIndex][sum] = 1
                return 1
            }
        }
        // skip the number and go to next index
        (*memo)[currentIndex][sum] = partRecursive(nums, sum, currentIndex + 1, memo)
    }
    return (*memo)[currentIndex][sum]

}

func PartEqualSumTab(nums []int) bool {
    // the idea behind tabulation is to create a sum from every subset
    canPart, sum := canPartition(nums)

    if !canPart {
        return false
    }

    tab := make([][]bool, len(nums))
    for i := 0; i < len(nums); i++ {
        tab[i] = make([]bool, sum + 1)
        for j := 0; j <= sum; j++ {
            tab[i][j] = false
        }
    }

    // for i := 0; i < len(nums); i++ {
    //     tab[i][0] = true // since 0 + 0 is 0, therefore first column will be true
    // }

    for s := 0; s <= sum; s++ {
        // true if the num is exactly equal to sum for 1st row
        tab[0][s] = nums[0] == s
    }

    for i := 1; i < len(nums); i++ {
        for s := 1; s <= sum; s++ {
            if tab[i - 1][s] {
                tab[i][s] = tab[i - 1][s]
            } else if nums[i] <= s {
                tab[i][s] = tab[i - 1][s - nums[i]]
            }
        }
    }
    return tab[len(nums) - 1][sum]
}

func KnapsackMemo(values []int, weights []int, capacity int) int {
    // values: [1, 6, 10, 16]
    // weights: [1, 2, 3, 5]
    // capacity: 6
    memoKnapsack := make([][]int, len(values)) // top-down, we fill the array as we find value
    for i := 0; i < len(values); i++ {
        memoKnapsack[i] = make([]int, capacity + 1)
        for j := 0; j <= capacity; j++ {
            memoKnapsack[i][j] = -1
        }
    }

    maxValue := knapsackRecursive(values, weights, capacity, 0, &memoKnapsack)
    return maxValue
}

func KnapsackTab(values []int, weights []int, capacity int) int {
   tabKnapsack := make([][]int, len(values)) // bottom-up, we fill the array pre-emptively
    for i := 0; i < len(values); i++ {
        tabKnapsack[i] = make([]int, capacity + 1)
        for j := 0; j <= capacity; j++ {
            tabKnapsack[i][j] = 0 // fill it zeroes starting out
        }
    }
    for c := 0; c <= capacity; c++ {
        if weights[0] <= c {
            tabKnapsack[0][c] = values[0] // fill the first row with the first value if weight is less
        }
    }

    for i := 1; i < len(values); i++ {
        for c := 1; c <= capacity; c++ {
            valueInclude := 0
            valueExclude := 0

            if weights[i] <= c {
                // include the item if weight is less than capacity
                valueInclude = values[i] + tabKnapsack[i - 1][c - weights[i]]
            }
            // exclude the item
            valueExclude = tabKnapsack[i - 1][c]
            if valueInclude >= valueExclude {
                tabKnapsack[i][c] = valueInclude
            } else {
                tabKnapsack[i][c] = valueExclude
            }
        }
    }
    fmt.Println(tabKnapsack)
    return tabKnapsack[len(values) - 1][capacity]
}

func knapsackRecursive(values []int, weights []int, capacity int, currentIndex int, memoKnapsack *[][]int) int {
    // base check
    if currentIndex >= len(values) || capacity <= 0 {
        return 0
    }

    if (*memoKnapsack)[currentIndex][capacity] != -1 {
        return (*memoKnapsack)[currentIndex][capacity]
    }

    value := values[currentIndex]
    weight := weights[currentIndex]

    valueInclude := 0
    valueExclude := 0

    if weight <= capacity {
        // if capacity is less than weight include the value
        // and increment the index while reducing capacity
        valueInclude = value + knapsackRecursive(values, weights, capacity - weight, currentIndex + 1, memoKnapsack)
    }
    // capacity is reached, skip this value
    valueExclude = knapsackRecursive(values, weights, capacity, currentIndex + 1, memoKnapsack)
    if (valueInclude >= valueExclude) {
        (*memoKnapsack)[currentIndex][capacity] = valueInclude
    } else {
        (*memoKnapsack)[currentIndex][capacity] = valueExclude
    }
    return (*memoKnapsack)[currentIndex][capacity]
}

// @TODO
func OnesAndZeroes(strs []string, m int, n int) int {
    // subset of strs that contain at most m 0's and n 1's
    // strs = ["10", "0001", "111001", "1", "0"]
    // m = 5, n = 3
    
    return -1
}
