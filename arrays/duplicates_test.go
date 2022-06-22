package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_containsDuplicates(t *testing.T) {
	type testCase struct {
		name string
		nums []int
		want bool
	}

	var tests = []testCase{
		{"example 1", []int{1, 2, 3, 1}, true},
		{"example 2", []int{1, 2, 3, 4}, false},
		{"example 3", []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := containsDuplicate(tc.nums)

			assert.Equal(t, tc.want, got)
		})
	}
}
