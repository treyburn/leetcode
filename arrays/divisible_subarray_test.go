package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_subarraysDivByK(t *testing.T) {
	type testCase struct {
		name      string
		input     []int
		divisible int
		want      int
	}

	var tests = []testCase{
		{"example 1", []int{4, 5, 0, -2, -3, 1}, 5, 7},
		{"no solution", []int{5}, 9, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := subarraysDivByK(tc.input, tc.divisible)

			assert.Equal(t, tc.want, got)
		})
	}
}
