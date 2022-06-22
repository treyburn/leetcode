package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_pivotIndex(t *testing.T) {
	type testCase struct {
		name string
		nums []int
		want int
	}

	var tests = []testCase{
		{"example 1", []int{1, 7, 3, 6, 5, 6}, 3},
		{"example 2", []int{1, 2, 3}, -1},
		{"example 3", []int{2, 1, -1}, 0},
		{"single item", []int{1}, 0},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := pivotIndex(tc.nums)

			assert.Equal(t, tc.want, got)
		})
	}
}
