package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_frequencySort(t *testing.T) {
	type testCase struct {
		name  string
		input []int
		want  []int
	}

	var tests = []testCase{
		{"example 1", []int{1, 1, 2, 2, 2, 3}, []int{3, 1, 1, 2, 2, 2}},
		{"exmaple 2", []int{2, 3, 1, 3, 2}, []int{1, 3, 3, 2, 2}},
		{"example 3", []int{-1, 1, -6, 4, 5, -6, 1, 4, 1}, []int{5, -1, 4, 4, -6, -6, 1, 1, 1}},
		{"example 4", []int{8, -8, 2, -8, -5, -3}, []int{8, 2, -3, -5, -8, -8}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := frequencySort(tc.input)

			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_descendingQuickSort(t *testing.T) {
	input := []int{8, 2, -5, -3}
	want := []int{8, 2, -3, -5}

	got := descendingQuickSort(input)

	assert.Equal(t, want, got)
}