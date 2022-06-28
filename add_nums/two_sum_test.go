package add_nums

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type testCase struct {
		name   string
		input  []int
		target int
		want   []int
	}

	var tests = []testCase{
		{"example 1", []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{"example 2", []int{3, 2, 4}, 6, []int{1, 2}},
		{"example 3", []int{3, 3}, 6, []int{0, 1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := twoSum(tc.input, tc.target)

			assert.Equal(t, tc.want, got)
		})
	}
}
