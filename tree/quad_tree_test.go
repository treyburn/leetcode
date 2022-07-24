package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	type testCase struct {
		name  string
		input [][]int
		want  *Node
	}

	var tests = []testCase{
		{"simple", [][]int{{0, 1}, {1, 0}}, &Node{
			Val:         true,
			IsLeaf:      false,
			TopLeft:     &Node{Val: false, IsLeaf: true},
			TopRight:    &Node{Val: true, IsLeaf: true},
			BottomLeft:  &Node{Val: true, IsLeaf: true},
			BottomRight: &Node{Val: false, IsLeaf: true},
		}},
		{"2 layer",
			[][]int{{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1, 1, 1, 1},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 0, 0, 0, 0}},
			&Node{
				Val:     true,
				IsLeaf:  false,
				TopLeft: &Node{Val: true, IsLeaf: true},
				TopRight: &Node{
					Val:         true,
					IsLeaf:      false,
					TopLeft:     &Node{Val: false, IsLeaf: true},
					TopRight:    &Node{Val: false, IsLeaf: true},
					BottomLeft:  &Node{Val: true, IsLeaf: true},
					BottomRight: &Node{Val: true, IsLeaf: true},
				},
				BottomLeft:  &Node{Val: true, IsLeaf: true},
				BottomRight: &Node{Val: false, IsLeaf: true},
			}},
		{"root is leaf", [][]int{{1, 1}, {1, 1}}, &Node{Val: true, IsLeaf: true}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := construct(tc.input)

			assert.Equal(t, tc.want, got)
		})
	}
}
