package substring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	type testCase struct {
		name  string
		input string
		want  string
	}

	var tests = []testCase{
		{"example1", "babad", "bab"},
		{"example2", "cbbd", "bb"},
		{"repeated1", "ababbaba", "ababbaba"},
		{"repeated2", "ababababa", "ababababa"},
		{"no palindrome", "12345", "1"},
		{"multiple palindromes", "abcbaabc", "cbaabc"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := longestPalindrome(tc.input)

			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_isPalindrome(t *testing.T) {
	type testCase struct {
		name  string
		input []rune
		want  bool
	}

	var tests = []testCase{
		{"basic palindrome", []rune("aba"), true},
		{"basic not a palindrome", []rune("abc"), false},
		{"repeated palindome", []rune("abababa"), true},
		{"repeated complex", []rune("abababbababa"), true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := isPalindrome(tc.input)

			assert.Equal(t, tc.want, got)
		})
	}
}
