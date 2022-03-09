package merge_intervals

import (
	"fmt"
	"sort"
)

func sortIntervals(intervals [][]int) {
    sort.Slice(intervals, func(i int, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
}

func MergeIntervals(input [][]int) [][]int {
    if (len(input) <= 1) {
        return input
    }
    mergedIntervals := make([][]int, 0)
    sortIntervals(input)
    start := input[0][0]
    end := input[0][1]

    for i := 1; i < len(input); i++ {
        interval := input[i]
        fmt.Println(start, end, interval)
        if end >= interval[0] { 
            fmt.Println("Found overlapping")
            // overlapping intervals
            // eg. 1, 3 and 2, 6
            if (end < interval[1]) {
                end = interval[1]
            } 
        } else {
            // non-overlapping intervals
            // eg. 8, 10 and 15, 18
            mergedIntervals = append(mergedIntervals, []int{start, end})
            start = interval[0]
            end = interval[1]
            
        }
    }
    mergedIntervals = append(mergedIntervals, []int{start, end})
    return mergedIntervals
}
