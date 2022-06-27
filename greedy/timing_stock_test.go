package greedy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxProfit(t *testing.T) {
	type testCase struct {
		name   string
		prices []int
		want   int
	}

	var tests = []testCase{
		{"example 1", []int{7, 1, 5, 3, 6, 4}, 5},
		{"example 2", []int{7, 6, 4, 3, 1}, 0},
		{"constant increase", []int{1, 2, 3, 4, 5, 6}, 5},
		{"simple 1", []int{1, 2}, 1},
		{"simple 2", []int{2, 1}, 0},
		{"1 item array", []int{1}, 0},
		{"simple 3", []int{2, 1, 4}, 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := maxProfit(tc.prices)

			assert.Equal(t, tc.want, got)
		})
	}
}
