package linked_lists

// naive solution O(2n)
func middleNode(head *ListNode) *ListNode {
	var length = 1
	var current = head

	for current.Next != nil {
		current = current.Next
		length++
	}

	current = head
	for i := 0; i < (length / 2); i++ {
		current = current.Next
	}
	return current
}

// fast and slow pointer solution
// roughly 35% faster
func middleNode2(head *ListNode) *ListNode {
	var slow = head
	var fast = head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
