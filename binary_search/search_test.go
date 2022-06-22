package binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_search(t *testing.T) {
	type testCase struct {
		name   string
		nums   []int
		target int
		want   int
	}

	var tests = []testCase{
		{"example 1", []int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{"example 2", []int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{"single item", []int{5}, 5, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := search(tc.nums, tc.target)

			assert.Equal(t, tc.want, got)
		})
	}
}
