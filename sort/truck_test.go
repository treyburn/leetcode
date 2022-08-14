package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maximumUnits(t *testing.T) {
	type testCase struct {
		name  string
		boxes [][]int
		size  int
		want  int
	}

	var tests = []testCase{
		{"example 1", [][]int{{1, 3}, {2, 2}, {3, 1}}, 4, 8},
		{"example 2", [][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, 10, 91},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := maximumUnits(tc.boxes, tc.size)

			assert.Equal(t, tc.want, got)
		})
	}
}
