package sll

import "fmt"

type Node struct {
	data int
	next *Node
}

func (node *Node) Next() *Node {
	return node.next
}

type SinglyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func (sll *SinglyLinkedList) Size() int {
	return sll.size
}

func (sll SinglyLinkedList) Print() {
	node := sll.head
	for node != nil {
		fmt.Printf("%v -> ", node.data)
		node = node.next
	}
	fmt.Println()
}

func (sll *SinglyLinkedList) GetHead() *Node {
	return sll.head
}

func (sll *SinglyLinkedList) Enqueue(data int) {
	node := &Node{data, nil}
	if sll.size == 0 {
		sll.head = node
		sll.tail = node
	} else {
		node.next = sll.head
		sll.head = node
	}
	sll.size++
}

func (sll *SinglyLinkedList) Dequeue() *Node {
	if sll.size == 0 {
		return nil
	} else {
		node := sll.head
		if sll.size == 1 {
			sll.head = nil
			sll.tail = nil
		} else {
			sll.head = sll.head.next
		}
		sll.size--
		return node
	}
}

func (sll *SinglyLinkedList) Push(data int) {
	node := &Node{data, nil}
	if sll.size == 0 {
		sll.head = node
		sll.tail = node
	} else {
		sll.tail.next = node
		sll.tail = node
	}
	sll.size++
}

func (sll *SinglyLinkedList) Pop() *Node {
	if sll.size == 0 {
		return nil
	} else {
		currentNode := sll.head
		secondLastNode := sll.head
		for currentNode.next != nil {
			secondLastNode = currentNode
			currentNode = currentNode.next
		}
		secondLastNode.next = nil
		sll.tail = secondLastNode
		sll.size--
		return currentNode
	}
}

func (sll *SinglyLinkedList) GetAt(index int) *Node {
	if sll.size == 0 || index < 0 || index >= sll.size {
		return nil
	} else {
		counter := 0
		currentNode := sll.head
		for counter < index {
			currentNode = currentNode.next
			counter++
		}
		return currentNode
	}
}

func (sll *SinglyLinkedList) InsertAt(index int, data int) {
	if index < 0 || index > sll.size {
		return
	} else {
		if index == 0 {
			// enqueue
			sll.Enqueue(data)
		} else if index == sll.size {
			// push
			sll.Push(data)
		} else {
			// retrieve previous node
			node := &Node{data, nil}
			prevNode := sll.GetAt(index - 1)
			nextNode := prevNode.next
			prevNode.next = node
			node.next = nextNode

			sll.size++
		}
	}
}

func (sll *SinglyLinkedList) RemoveAt(index int) *Node {
	if index < 0 || index >= sll.size {
		return nil
	} else {
		if index == 0 {
			return sll.Dequeue()
		} else if index == sll.size-1 {
			return sll.Pop()
		} else {
			prevNode := sll.GetAt(index - 1)
			nodeToRemove := prevNode.next
			nextNode := nodeToRemove.next
			prevNode.next = nextNode
			nodeToRemove.next = nil

			sll.size--
			return nodeToRemove
		}
	}
}

// In place reversal of linked list
func (sll *SinglyLinkedList) Reverse() {
	currentNode := sll.head
	var nextNode *Node = nil
	var prevNode *Node = nil
	for currentNode != nil {
		nextNode = currentNode.next
		currentNode.next = prevNode
		prevNode = currentNode
		currentNode = nextNode
	}
	sll.tail = sll.head
	sll.head = prevNode
	fmt.Print("Reversed list: ")
	sll.Print()
}

// Connect tail to the node at index
func (sll *SinglyLinkedList) Connect(index int) {
	if index < 0 || index >= sll.size-1 {
		return
	} else {
		nodeToConnect := sll.GetAt(index)
		sll.tail.next = nodeToConnect
	}
}
