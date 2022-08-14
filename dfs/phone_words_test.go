package dfs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	type testCase struct {
		name    string
		input   string
		wantLen int
	}

	var tests = []testCase{
		{"example 1", "23", 9},
		{"example 2", "", 0},
		{"example 3", "2", 3},
		{"example 4", "239", 36},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := letterCombinations(tc.input)

			assert.Equal(t, tc.wantLen, len(got))
		})
	}
}
