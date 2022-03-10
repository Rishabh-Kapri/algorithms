// Heap is tree DS in which the tree is a complete binary tree
// A complete binary tree is a binary tree in which all the levels are
// completely filled excepty possibly the lowest one which is filled from the
// left

// Min Heap
// If each parent node is less than or equal to its child node
// Max Heap
// If each parent node is greater than or equal to its child node

// Root node, i = 0
// A parent node, parent(i) = (i - 1) / 2
// A left child node, left(i) = 2*i + 1
// A right child node, right(i) = 2*i + 2

// Time complexity
// Search: O(n)
// Insert: O(log n)
// Delete: O(log n)
// Peek: O(1)

package heap

type Heap struct {
	arr         []int
	compareFunc func(int, int) bool
}

// Init a heap
func Constructor(arr []int, compareFunc func(int, int) bool) Heap {
	return Heap{
		arr,
		compareFunc,
	}
}

// Returns the value at the top of the heap
func (heap *Heap) Peek() int {
	return heap.arr[0]
}

// Returns the whole heap
func (heap *Heap) Heap() []int {
	return heap.arr
}

// Value at an index
func (heap *Heap) Val(index int) int {
	if index < heap.Size() && index >= 0 {
		return heap.arr[index]
	}
	return -1
}

func (heap *Heap) Size() int {
	return len(heap.arr)
}

// Swap values at two index
func (heap *Heap) Swap(index1 int, index2 int) {
	if index1 < heap.Size() && index1 >= 0 && index2 < heap.Size() && index2 >= 0 {
		heap.arr[index1], heap.arr[index2] = heap.arr[index2], heap.arr[index1]
	}
}

func (heap *Heap) getLeftChild(index int) int {
	return 2*index + 1
}

func (heap *Heap) getRightChild(index int) int {
	return 2*index + 2
}

func (heap *Heap) getParent(index int) int {
	return (index - 1) / 2
}

func (heap *Heap) heapifyTopDown(index int) {
	left := heap.getLeftChild(index)
	right := heap.getRightChild(index)
	i := index

	if left < heap.Size() && heap.compareFunc(heap.Val(i), heap.Val(left)) {
		i = left
	}
	if right < heap.Size() && heap.compareFunc(heap.Val(i), heap.Val(right)) {
		i = right
	}

	if i != index {
		heap.Swap(index, i)
		heap.heapifyTopDown(i)
	} else {

	}
}

func (heap *Heap) heapifyBottomUp(index int) {
	parent := heap.getParent(index)

	val := heap.Val(index)
	parentVal := heap.Val(parent)

	if heap.Size() <= 1 {
		return
	}

	if heap.compareFunc(parentVal, val) {
		heap.Swap(parent, index)
		heap.heapifyBottomUp(parent)
	}
}

// Push an element to the heap
func (heap *Heap) Push(num int) {
	heap.arr = append(heap.arr, num)
	heap.heapifyBottomUp(heap.Size() - 1)
}

// Pop the first element from the heap
func (heap *Heap) Pop() int {
	ele := heap.arr[0]
	// put last element at the start to make sure the top down heapify function
	// is called recursively for child elements too
	// because if 1st element is already smaller then no other recursive calls
	// will happen
	heap.arr[0] = heap.arr[heap.Size()-1]
	heap.arr = heap.arr[:heap.Size()-1]
	heap.heapifyTopDown(0)
	return ele
}

// Buld the heap using the array and compare function
func (heap *Heap) BuildHeap() {
	for i := (heap.Size() - 1); i >= 0; i-- {
		heap.heapifyBottomUp(i)
	}
}
