package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_myPow(t *testing.T) {
	type testCase struct {
		name string
		x    float64
		n    int
		want float64
	}

	var tests = []testCase{
		{"simple", 2.00000, 10, 1024.00000},
		{"decimal", 2.10000, 3, 9.26100},
		{"negative pow", 2.00000, -2, 0.25000},
		{"zero", 5.00000, 0, 1.00000},
		{"one", 500.10000, 1, 500.10000},
		{"long decimal", 8.88023, 3, 700.28148},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := myPow(tc.x, tc.n)

			assert.Equal(t, tc.want, got)
		})
	}
}
