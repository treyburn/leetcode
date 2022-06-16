package substring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	name  string
	input string
	want  int
}

var tests = []testCase{
	{"example 1", "abcabcbb", 3},
	{"example 2", "bbbbb", 1},
	{"example 3", "pwwkew", 3},
	{"space", " ", 1},
	{"backtrack", "dvdf", 3},
	{"else", "ckilbkd", 5},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := lengthOfLongestSubstring(tc.input)
			assert.Equal(t, tc.want, got)

		})
	}
}
