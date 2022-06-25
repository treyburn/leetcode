package linked_lists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type testCase struct {
		name  string
		list1 *ListNode
		list2 *ListNode
		want  *ListNode
	}

	var tests = []testCase{
		{"example 1", listNodeFromSlice([]int{1, 2, 4}), listNodeFromSlice([]int{1, 3, 4}), listNodeFromSlice([]int{1, 1, 2, 3, 4, 4})},
		{"example 2", listNodeFromSlice([]int{}), listNodeFromSlice([]int{}), listNodeFromSlice([]int{})},
		{"example 3", listNodeFromSlice([]int{1}), listNodeFromSlice([]int{}), listNodeFromSlice([]int{1})},
		{"example 4", listNodeFromSlice([]int{}), listNodeFromSlice([]int{1}), listNodeFromSlice([]int{1})},
		{"example 5", listNodeFromSlice([]int{2}), listNodeFromSlice([]int{1}), listNodeFromSlice([]int{1, 2})},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := mergeTwoLists(tc.list1, tc.list2)

			assert.Equal(t, tc.want, got)
		})
	}

}

func Test_mergeTwoLists_modifed(t *testing.T) {
	type testCase struct {
		name  string
		list1 *ListNode
		list2 *ListNode
		want  *ListNode
	}

	var tests = []testCase{
		{"example 1", listNodeFromSlice([]int{1, 2, 4}), listNodeFromSlice([]int{1, 3, 4}), listNodeFromSlice([]int{1, 1, 2, 3, 4, 4})},
		{"example 2", listNodeFromSlice([]int{}), listNodeFromSlice([]int{}), listNodeFromSlice([]int{})},
		{"example 3", listNodeFromSlice([]int{1}), listNodeFromSlice([]int{}), listNodeFromSlice([]int{1})},
		{"example 4", listNodeFromSlice([]int{}), listNodeFromSlice([]int{1}), listNodeFromSlice([]int{1})},
		{"example 5", listNodeFromSlice([]int{2}), listNodeFromSlice([]int{1}), listNodeFromSlice([]int{1, 2})},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := mergeTwoLists_modified(tc.list1, tc.list2)

			assert.Equal(t, tc.want, got)
		})
	}

}

func listNodeFromSlice(n []int) *ListNode {
	if len(n) == 0 {
		return nil
	}
	return &ListNode{
		Val:  n[0],
		Next: listNodeFromSlice(n[1:]),
	}
}

var res *ListNode
var l1 = listNodeFromSlice([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 11, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 11, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3})
var l2 = listNodeFromSlice([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 11, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 11, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3})

func Benchmark_original_merge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res = mergeTwoLists(l1, l2)
	}
}

func Benchmark_modified_merge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res = mergeTwoLists_modified(l1, l2)
	}
}

func Benchmark_original_merge_best_case(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res = mergeTwoLists(nil, l2)
	}
}

func Benchmark_modified_merge_best_case(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res = mergeTwoLists_modified(nil, l2)
	}
}
