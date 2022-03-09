package fast_slow_pointers

import sll "algo_go/singly_linked_list"

// returns true if cycle exists
func LinkedListCycleEasy(head *sll.Node) bool {
	slow := head
	fast := head

	// keep the loop running
	for fast != nil && fast.Next() != nil {
		slow = slow.Next()
		fast = fast.Next().Next()

		if slow == fast {
			return true
		}
	}
	return false
}

// returns the node at which the cycle exists
func LinkedListCycleMedium(head *sll.Node) *sll.Node {
	slow := head
	fast := head

	// keep the loop running
	for fast != nil && fast.Next() != nil {
		slow = slow.Next()
		fast = fast.Next().Next()

		if slow == fast {
			current := head
			for current != slow {
				current = current.Next()
				slow = slow.Next()
			}
			return slow
		}
	}
	return nil
}
