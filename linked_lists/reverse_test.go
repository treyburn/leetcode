package linked_lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type testCase struct {
		name string
		list *ListNode
		want *ListNode
	}

	var tests = []testCase{
		{"example 1", listNodeFromSlice([]int{1, 2, 3, 4, 5}), listNodeFromSlice([]int{5, 4, 3, 2, 1})},
		{"example 2", listNodeFromSlice([]int{1, 2}), listNodeFromSlice([]int{2, 1})},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := reverseList(tc.list)

			assert.Equal(t, tc.want, got)
		})
	}
}
