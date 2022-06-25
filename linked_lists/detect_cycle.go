package linked_lists

// O(n) time complexity and O(n) space complexity
func detectCycle(head *ListNode) *ListNode {
	var set = make(map[*ListNode]struct{}, 0)
	var current = head

	for current != nil && current.Next != nil {
		_, ok := set[current]
		if !ok {
			set[current] = struct{}{}
			current = current.Next
		} else {
			return current
		}
	}

	return nil
}
