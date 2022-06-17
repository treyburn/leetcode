package zigzag

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_convert(t *testing.T) {
	type testCase struct {
		name     string
		input    string
		rowCount int
		want     string
	}

	var tests = []testCase{
		{"example1", "PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"example2", "PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"example3", "A", 1, "A"},
		{"5 rows", "PAYPALISHIRING", 5, "PHASIYIRPLIGAN"},
		{"weird one", "AB", 1, "AB"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := convert(tc.input, tc.rowCount)

			assert.Equal(t, tc.want, got)
		})
	}
}
