package linked_lists

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tail := reverseList(head.Next)
	head.Next.Next = head // head.next.next == tail.next on the first iteration, then tail.next.next, then tail.next.next.next
	head.Next = nil       // set to nil to avoid a circular list
	return tail
}
