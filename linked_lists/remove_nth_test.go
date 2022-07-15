package linked_lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_removeNth(t *testing.T) {
	input := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	want := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
	}

	got := removeNthFromEnd2(input, 2)

	assert.Equal(t, want, got)
}

func Test_removeNth_one(t *testing.T) {
	input := &ListNode{
		Val:  1,
		Next: nil,
	}

	got := removeNthFromEnd2(input, 1)

	assert.Nil(t, got)
}

func Test_removeNth_two(t *testing.T) {
	input := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}

	want := &ListNode{
		Val:  1,
		Next: nil,
	}

	got := removeNthFromEnd2(input, 1)

	assert.Equal(t, want, got)
}

func Test_removeNth_two2(t *testing.T) {
	input := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}

	want := &ListNode{
		Val:  2,
		Next: nil,
	}

	got := removeNthFromEnd2(input, 2)

	assert.Equal(t, want, got)
}
