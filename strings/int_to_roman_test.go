package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_intToRoman(t *testing.T) {
	type testCase struct {
		name  string
		input int
		want  string
	}

	var tests = []testCase{
		{"one", 1, "I"},
		{"three", 3, "III"},
		{"four", 4, "IV"},
		{"five", 5, "V"},
		{"nine", 9, "IX"},
		{"forty", 40, "XL"},
		{"fifty-eight", 58, "LVIII"},
		{"ninety-nine", 99, "XCIX"},
		{"1984", 1994, "MCMXCIV"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := intToRoman(tc.input)

			assert.Equal(t, tc.want, got)
		})
	}
}
