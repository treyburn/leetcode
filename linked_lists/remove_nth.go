package linked_lists

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	cache := []*ListNode{head}
	next := head.Next

	for next != nil {
		cache = append(cache, next)
		next = next.Next
	}

	if len(cache)-n == 0 {
		return head.Next
	}

	prev := cache[len(cache)-(n+1)]
	if prev.Next == nil {
		return head
	} else {
		prev.Next = prev.Next.Next
	}

	return head
}

// a better solution is to use 2 pointers instead of 1 pointer + array
// advance first pointer by n+1 and then both pointers are n units apart. advance both to next until first pointer is nil
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	firstPtr := head
	secondPtr := head

	for i := 0; i < n; i++ {
		secondPtr = secondPtr.Next
	}

	for secondPtr != nil && secondPtr.Next != nil {
		firstPtr = firstPtr.Next
		secondPtr = secondPtr.Next
	}

	if secondPtr == nil {
		return firstPtr.Next
	}

	firstPtr.Next = firstPtr.Next.Next

	return head
}
