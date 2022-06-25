package linked_lists

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	switch {
	case list1 == nil && list2 == nil:
		return nil
	case list1 == nil && list2 != nil:
		return &ListNode{
			Val:  list2.Val,
			Next: mergeTwoLists(nil, list2.Next),
		}
	case list1 != nil && list2 == nil:
		return &ListNode{
			Val:  list1.Val,
			Next: mergeTwoLists(list1.Next, nil),
		}
	case list1.Val <= list2.Val:
		return &ListNode{
			Val:  list1.Val,
			Next: mergeTwoLists(list1.Next, list2),
		}
	default:
		return &ListNode{
			Val:  list2.Val,
			Next: mergeTwoLists(list1, list2.Next),
		}
	}
}

// in a best case scenario (when one of the first lists is already nil) - the modified solution is substantially faster by ceasing to advance and just returning the remainder
func mergeTwoLists_modified(list1 *ListNode, list2 *ListNode) *ListNode {
	switch {
	case list1 == nil && list2 == nil:
		return nil
	case list1 == nil && list2 != nil:
		return list2
	case list1 != nil && list2 == nil:
		return list1
	case list1.Val <= list2.Val:
		return &ListNode{
			Val:  list1.Val,
			Next: mergeTwoLists_modified(list1.Next, list2),
		}
	default:
		return &ListNode{
			Val:  list2.Val,
			Next: mergeTwoLists_modified(list1, list2.Next),
		}
	}
}
