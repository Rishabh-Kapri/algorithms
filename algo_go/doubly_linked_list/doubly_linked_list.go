package dll

import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func (dll *DoublyLinkedList) Size() int {
	return dll.size
}

func (dll DoublyLinkedList) Print() {
	node := dll.head
	for node != nil {
		str := ""
		if node.prev == nil {
			str += "nil | "
		} else {
			str += fmt.Sprint(node.prev.data) + "   | "
		}
		str += fmt.Sprint(node.data)
		if node.next == nil {
			str += " | nil"
		} else {
			str += " |  " + fmt.Sprint(node.next.data)
		}
		fmt.Println(str)
		node = node.next
	}
	fmt.Println()
}

func (dll *DoublyLinkedList) Enqueue(data int) {
	node := &Node{data, nil, nil}
	if dll.size == 0 {
		dll.head = node
		dll.tail = node
	} else {
		dll.head.prev = node
		node.next = dll.head
		dll.head = node
	}
	dll.size++
}

func (dll *DoublyLinkedList) Dequeue() *Node {
	if dll.size == 0 {
		return nil
	} else {
		node := dll.head
		if dll.size == 1 {
			dll.head = nil
			dll.tail = nil
		} else {
			dll.head = dll.head.next
			dll.head.prev = nil
		}
		dll.size--
		return node
	}
}

func (dll *DoublyLinkedList) Push(data int) {
	node := &Node{data, nil, nil}
	if dll.size == 0 {
		dll.head = node
		dll.tail = node
	} else {
		dll.tail.next = node
		node.prev = dll.tail
		dll.tail = node
	}
	dll.size++
}

func (dll *DoublyLinkedList) Pop() *Node {
	if dll.size == 0 {
		return nil
	} else {
		node := dll.tail
		if dll.size == 1 {
			dll.head = nil
			dll.tail = nil
		} else {
			dll.tail = dll.tail.prev
			dll.tail.next = nil
		}
		dll.size--
		return node
	}
}

func (dll *DoublyLinkedList) GetAt(index int) *Node {
	if dll.size == 0 || index < 0 || index >= dll.size {
		return nil
	} else {
		counter := 0
		currentNode := dll.head
		for counter < index {
			currentNode = currentNode.next
			counter++
		}
		return currentNode
	}
}

func (dll *DoublyLinkedList) InsertAt(index int, data int) {
	if index < 0 || index > dll.size {
		return
	} else {
		if index == 0 {
			// enqueue
			dll.Enqueue(data)
		} else if index == dll.size {
			// push
			dll.Push(data)
		} else {
			node := &Node{data, nil, nil}
			prevNode := dll.GetAt(index - 1)
			nextNode := prevNode.next
			prevNode.next = node
			node.prev = prevNode
			node.next = nextNode
			nextNode.prev = node

			dll.size++
		}
	}
}

func (dll *DoublyLinkedList) RemoveAt(index int) *Node {
	if index < 0 || index >= dll.size {
		return nil
	} else {
		if index == 0 {
			// dequeue
			return dll.Dequeue()
		} else if index == dll.size-1 {
			// pop
			return dll.Pop()
		} else {
			nodeToRemove := dll.GetAt(index)
			prevNode := nodeToRemove.prev
			nextNode := nodeToRemove.next
			prevNode.next = nextNode
			nextNode.prev = prevNode
			nodeToRemove.next = nil
			nodeToRemove.prev = nil

			dll.size--
			return nodeToRemove

		}
	}
}
