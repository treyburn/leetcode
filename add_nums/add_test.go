package add_nums

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddTwoNums(t *testing.T) {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	want := &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}}

	got := AddTwoNumbers(l1, l2)

	assert.Equal(t, want, got)
}
