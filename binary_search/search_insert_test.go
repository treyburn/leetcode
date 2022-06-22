package binary_search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_searchInsert(t *testing.T) {
	type testCase struct {
		name   string
		nums   []int
		target int
		want   int
	}

	var tests = []testCase{
		{"example 1", []int{1, 3, 5, 6}, 5, 2},
		{"example 2", []int{1, 3, 5, 6}, 2, 1},
		{"example 3", []int{1, 3, 5, 6}, 7, 4},
		{"example 4", []int{1, 3, 5, 6}, 0, 0},
		{"example 5", []int{1, 3}, 2, 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := searchInsert(tc.nums, tc.target)

			assert.Equal(t, tc.want, got)
		})
	}
}
