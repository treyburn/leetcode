package narray_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var example2 = &Node{
	Val: 1,
	Children: []*Node{
		{Val: 2},
		{Val: 3, Children: []*Node{
			{Val: 6},
			{Val: 7, Children: []*Node{
				{Val: 11, Children: []*Node{
					{Val: 14},
				}},
			}},
		}},
		{Val: 4, Children: []*Node{
			{Val: 8, Children: []*Node{
				{Val: 12},
			}},
		}},
		{Val: 5, Children: []*Node{
			{Val: 9, Children: []*Node{
				{Val: 13},
			}},
			{Val: 10},
		}},
	},
}

func Test_preorder(t *testing.T) {
	type testCase struct {
		name  string
		input *Node
		want  []int
	}

	example1 := &Node{
		Val: 1,
		Children: []*Node{
			{
				Val:      3,
				Children: []*Node{{Val: 5}, {Val: 6}},
			},
			{Val: 2},
			{Val: 4},
		},
	}

	var tests = []testCase{
		{"example 1", example1, []int{1, 3, 5, 6, 2, 4}},
		{"example 2", example2, []int{1, 2, 3, 6, 7, 11, 14, 4, 8, 12, 5, 9, 13, 10}},
		{"nil", nil, nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := preorder(tc.input)

			assert.Equal(t, tc.want, got)
		})
	}
}

var preoderResult []int

func Benchmark_preorder_recursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		preoderResult = preorder_recrusive(example2)
	}
}

func Benchmark_preorder_iterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		preoderResult = preorder(example2)
	}
}
