package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_runningSum(t *testing.T) {
	type testCase struct {
		name string
		nums []int
		want []int
	}

	var tests = []testCase{
		{"example 1", []int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
		{"example 2", []int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
		{"example 3", []int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := runningSum(tc.nums)

			assert.Equal(t, tc.want, got)
		})
	}
}
