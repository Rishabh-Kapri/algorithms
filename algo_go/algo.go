package main

import (
    "fmt"
    sll "algo_go/singly_linked_list"
    dll "algo_go/doubly_linked_list"
    "algo_go/two_pointers"
    "algo_go/sliding_window"
    "algo_go/fast_and_slow_pointers"
    "algo_go/merge_intervals"
    "algo_go/cyclic_sort"
    "algo_go/subsets"
    "algo_go/modified_binary_search"
    "algo_go/dp"
    "algo_go/heap"
    "algo_go/two_heaps"
)

func main() {
    sll1 := sll.SinglyLinkedList{}
    sll1.Enqueue(1)
    sll1.Enqueue(2)
    // sll1.Dequeue()
    sll1.Push(5)
    sll1.Push(3) 
    sll1.Push(7)
    sll1.Pop()
    sll1.Print()
    sll1.Reverse()

    dll := dll.DoublyLinkedList{}
    dll.Enqueue(3)
    dll.Enqueue(2)
    dll.Enqueue(1)
    dll.Dequeue()
    dll.Push(4)
    dll.Enqueue(0)
    dll.InsertAt(1, 1)
    dll.Print()
    fmt.Println(dll)

    heapMin := heap.Constructor(
        []int{5, 9, 11, 14, 18, 19, 21, 33, 17, 27}, 
        func(a int, b int) bool { return a > b },
    )
    heapMax := heap.Constructor(
        []int{5, 9, 11, 14, 18, 19, 21, 33, 17, 27}, 
        func(a int, b int) bool { return a < b },
    )
    heapMin.BuildHeap()
    fmt.Println("Min Heap:", heapMin.Heap())
    heapMax.BuildHeap()
    fmt.Println("Max Heap:", heapMax.Heap())

    a, b := two_pointers.TwoSum([]int{2, 7, 11, 15}, 9)
    fmt.Println("Two Pointers solution:", a, b)

    fmt.Println(sliding_window.MaxAvg([]int{1, 12, -5, -6, 50, 3}, 4))

    sll2 := sll.SinglyLinkedList{}
    sll2.Push(3)
    sll2.Push(2)
    sll2.Push(0)
    sll2.Push(-4)
    sll2.Push(1)
    sll2.Connect(2)
    // linked list: 3 -> 2 -> 0 -> -> 4 -v 
    //                   ^---------------
    fmt.Println("Loop Exist:", fast_slow_pointers.LinkedListCycleEasy(sll2.GetHead()))
    fmt.Println("Loop starts at node:", fast_slow_pointers.LinkedListCycleMedium(sll2.GetHead()))
    fmt.Println("Merged Intervals:", merge_intervals.MergeIntervals([][]int{
        {1, 4},
        {0, 4}, 
        {6, 10}, 
        {15, 18},
    }))
    fmt.Println("Missing number:", cyclic_sort.MissingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
    fmt.Println("Missing number:", cyclic_sort.MissingNumber([]int{0, 1}))
    fmt.Println("Subsets:", subsets.Subsets([]int{1, 2, 3}))
    fmt.Println("Num is at index:", binary_search.BinarySearch([]int{-1, 0 , 3, 5, 9, 12}, 9))
    fmt.Println("Num is at index:", binary_search.BinarySearch([]int{-1, 0 , 3, 5, 9, 12}, 2))

    fmt.Println("Fibonacci Memoization:", dp.FibonacciMemo(15))
    fmt.Println("Fibonacci Tabular:", dp.FibonacciTab(15))
    fmt.Println("Can partition memo:", dp.PartEqualSumMemo([]int{1, 5, 13, 5}))
    fmt.Println("Can partition tab:", dp.PartEqualSumTab([]int{1, 2, 3, 4, 5, 6, 7}))
    fmt.Println("Max knapsack value using memo:", dp.KnapsackMemo([]int{1, 6, 10, 16}, []int{1, 2, 3, 5}, 6))
    fmt.Println("Max knapsack value using tab:", dp.KnapsackTab([]int{1, 6, 10, 16}, []int{1, 2, 3, 5}, 6))

    twoHeap := two_heaps.Constructor()
    twoHeap.AddNum(6)
    fmt.Println("Median:", twoHeap.FindMedian())
    twoHeap.AddNum(10)
    fmt.Println("Median:", twoHeap.FindMedian())
    twoHeap.AddNum(2)
    fmt.Println("Median:", twoHeap.FindMedian())
    twoHeap.AddNum(6)
    fmt.Println("Median:", twoHeap.FindMedian())
    twoHeap.AddNum(5)
    fmt.Println("Median:", twoHeap.FindMedian())
    twoHeap.Print()
}


