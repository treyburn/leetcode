package linked_lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	type testCase struct {
		name string
		head *ListNode
		want *ListNode
	}

	circ1 := &ListNode{Val: 2, Next: &ListNode{Next: &ListNode{Val: -4, Next: nil}}}
	circ1.Next.Next.Next = circ1
	circ2 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}
	circ2.Next.Next = circ2

	var tests = []testCase{
		{"example 1", &ListNode{Val: 3, Next: circ1}, circ1},
		{"example 2", circ2, circ2},
		{"no cycle", listNodeFromSlice([]int{1, 2, 3}), nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := detectCycle(tc.head)

			assert.Equal(t, tc.want, got)
		})
	}
}
