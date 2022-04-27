package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode() *ListNode {
	return &ListNode{Next: &ListNode{}}
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

//l1 := [2,4,3]
//l2 := [5,6,4]

//Output: [7,0,8] [7, 0, 4]

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6}}

	result := addTwoNumbers(l1, l2)
	fmt.Println(result)
}
