package binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type testCase struct {
		name  string
		input *TreeNode
		want  [][]int
	}

	var example1 = &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	var example2 = &TreeNode{Val: 1}

	var tests = []testCase{
		{"example 1", example1, [][]int{{3}, {9, 20}, {15, 7}}},
		{"example 2", example2, [][]int{{1}}},
		{"example 3", nil, nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := levelOrder(tc.input)

			assert.Equal(t, tc.want, got)
		})
	}
}
