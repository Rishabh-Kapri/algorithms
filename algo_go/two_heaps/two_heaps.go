package two_heaps

import (
    "fmt"
    "algo_go/heap"
)

type MedianFinder struct { 
    minHeap *heap.Heap
    maxHeap *heap.Heap
}

func Constructor() MedianFinder {
    minHeap := heap.Constructor(
        []int{},
        func (a int, b int) bool {
            return a > b
        },
    );
    maxHeap := heap.Constructor(
        []int{},
        func (a int, b int) bool {
            return a < b
        },
    );
    return MedianFinder{
        minHeap: &minHeap,
        maxHeap: &maxHeap,
    }
}

func (this MedianFinder) Print() {
    fmt.Println("Max heap:", this.maxHeap.Heap(), "Min heap:", this.minHeap.Heap())
}

// left of median (containing smaller values) will be max heap, so we get largest value on top
// right of median (containing greater values) will be min heap, so we get smallest value on top
// median will be sum of top values of min and max heap if even number of elements
// for odd number of elements median will be top of max heap
func (this *MedianFinder) AddNum(num int) {
    if this.maxHeap.Size() == 0 || this.maxHeap.Peek() >= num {
        this.maxHeap.Push(num)
    } else {
        this.minHeap.Push(num)
    }

    // balance both heaps to have equal elements
    // in case of odd number of elements, put more in max heap
    if this.maxHeap.Size() > this.minHeap.Size() + 1 {
        // if max heap has an extra element, put top element to min heap
        this.minHeap.Push(this.maxHeap.Pop())
    } else if this.maxHeap.Size() < this.minHeap.Size() {
        // if min heap has more elements, put top element to max heap
        this.maxHeap.Push(this.minHeap.Pop())
    }
}

func (this *MedianFinder) FindMedian() float64 {
    if this.maxHeap.Size() == this.minHeap.Size() {
        // even number of elements
        return float64(float64(this.maxHeap.Peek() + this.minHeap.Peek()) / 2.0)
    } 
    // odd number of elements
    return float64(this.maxHeap.Peek())
}
