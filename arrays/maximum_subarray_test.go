package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	type testCase struct {
		name string
		nums []int
		want int
	}

	var tests = []testCase{
		{"example 1", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{"example 2", []int{1}, 1},
		{"example 3", []int{5, 4, -1, 7, 8}, 23},
		{"tricky example", []int{-1}, -1},
		{"example 4", []int{1, 2, -1, -2, 2, 1, -2, 1, 4, -5, 4}, 6},
		{"one big with lots of smalls", []int{2, 2, 2, 2, 2, 2, 2, -10, 5}, 14},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := maxSubArray(tc.nums)

			assert.Equal(t, tc.want, got)
		})
	}
}
