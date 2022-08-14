package chess

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_knightDialer(t *testing.T) {
	type testCase struct {
		name string
		n    int
		want int
	}

	var tests = []testCase{
		{"example 1", 1, 10},
		{"example 2", 2, 20},
		{"example 3", 3131, 136006598},
		{"example 4", 3353, 283772296},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := knightDialer(tc.n)

			assert.Equal(t, tc.want, got)
		})
	}
}
