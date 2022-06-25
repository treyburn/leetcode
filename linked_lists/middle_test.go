package linked_lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_middleNode(t *testing.T) {
	type testCase struct {
		name string
		head *ListNode
		want *ListNode
	}

	var tests = []testCase{
		{"example 1", listNodeFromSlice([]int{1, 2, 3, 4, 5}), listNodeFromSlice([]int{3, 4, 5})},
		{"example 2", listNodeFromSlice([]int{1, 2, 3, 4, 5, 6}), listNodeFromSlice([]int{4, 5, 6})},
		{"len of 1", listNodeFromSlice([]int{1}), listNodeFromSlice([]int{1})},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := middleNode(tc.head)

			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_middleNode2(t *testing.T) {
	type testCase struct {
		name string
		head *ListNode
		want *ListNode
	}

	var tests = []testCase{
		{"example 1", listNodeFromSlice([]int{1, 2, 3, 4, 5}), listNodeFromSlice([]int{3, 4, 5})},
		{"example 2", listNodeFromSlice([]int{1, 2, 3, 4, 5, 6}), listNodeFromSlice([]int{4, 5, 6})},
		{"len of 1", listNodeFromSlice([]int{1}), listNodeFromSlice([]int{1})},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := middleNode2(tc.head)

			assert.Equal(t, tc.want, got)
		})
	}
}

var resMiddleNode *ListNode
var list = listNodeFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9})

func Benchmark_middleNode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resMiddleNode = middleNode(list)
	}
}

func Benchmark_middleNode2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resMiddleNode = middleNode2(list)
	}
}
