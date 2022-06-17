package atoi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_myAtoi(t *testing.T) {
	type testCase struct {
		name  string
		input string
		want  int
	}

	var tests = []testCase{
		{"example 1", "42", 42},
		{"example 2", "-42", -42},
		{"bigger num", "4193", 4193},
		{"example 3", "4193 with words", 4193},
		{"too big - negative", "-9992147483648", -2147483648},
		{"too big - positive", "9992147483648", 2147483647},
		{"too big - positive 2", "9223372036854775808", 2147483647},
		{"spaces", "   -42", -42},
		{"break if words come first", "words and 987", 0},
		{"multiple modifiers", "+-12", 0},
		{"garbage", "00000-42a1234", 0},
		{"leading zeros", "  0000000000012345678", 12345678},
		{"more spaces", "  +  413", 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := myAtoi(tc.input)

			assert.Equal(t, tc.want, got)
		})
	}
}
