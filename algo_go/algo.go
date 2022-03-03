package main

import (
    sll "algo_go/singly_linked_list"
    dll "algo_go/doubly_linked_list"
    "fmt"
)

func main() {
    sll := sll.SinglyLinkedList{}
    sll.Enqueue(1)
    sll.Enqueue(2)
    // sll.Dequeue()
    sll.Push(5)
    sll.Push(3)
    sll.Push(7)
    sll.Pop()
    sll.Print()
    sll.Reverse()

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
}


