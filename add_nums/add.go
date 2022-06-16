package add_nums

// ListNode is a singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var remainder int
	first := true
	var result, current *ListNode
	for l1 != nil || l2 != nil || remainder != 0 {
		var n1, n2 int
		if !(l1 != nil) {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}
		if !(l2 != nil) {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}
		v := n1 + n2 + remainder
		if v >= 10 {
			remainder = 1
			v -= 10
		} else {
			remainder = 0
		}
		if first {
			first = false
			result = &ListNode{}
			current = result
		} else {
			current.Next = &ListNode{}
			current = current.Next
		}
		current.Val = v
	}
	return result
}
