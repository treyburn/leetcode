package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type testCase struct {
		name  string
		input string
		want  int
	}

	var tests = []testCase{
		{"example 1", "abccccdd", 7},
		{"example 2", "a", 1},
		{"example 3", "ccc", 3},
		{"example 4", "cccc", 4},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := longestPalindrome(tc.input)

			assert.Equal(t, tc.want, got)
		})
	}
}
