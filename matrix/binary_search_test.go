package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_leftMostColumnWithOne(t *testing.T) {
	type testCase struct {
		name  string
		input [][]int
		want  int
	}

	var tests = []testCase{
		{"simple", [][]int{{0, 0}, {1, 1}}, 0},
		{"simple2", [][]int{{0, 0}, {0, 1}}, 1},
		{"none", [][]int{{0, 0}, {0, 0}}, -1},
		{"uneven", [][]int{
			{0, 0, 0, 0, 0},
			{0, 0, 1, 1, 1},
			{0, 0, 0, 1, 1},
			{0, 1, 1, 1, 1},
			{0, 0, 1, 1, 1},
			{0, 0, 0, 0, 1},
		}, 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			bm := makeBinaryMatrix(tc.input)

			got := leftMostColumnWithOne(bm)

			assert.Equal(t, tc.want, got)
		})
	}
}

func makeBinaryMatrix(input [][]int) BinaryMatrix {
	return BinaryMatrix{
		Get: func(x int, y int) int {
			return input[x][y]
		},
		Dimensions: func() []int {
			return []int{len(input), len(input[0])}
		},
	}
}
