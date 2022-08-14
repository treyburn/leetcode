package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findOrder(t *testing.T) {
	type testCase struct {
		name       string
		prereqs    [][]int
		numCourses int
		want       []int
	}

	var tests = []testCase{
		{"example 1", [][]int{{1, 0}}, 2, []int{0, 1}},
		{"example 2", [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, 4, []int{0, 1, 2, 3}},
		{"example 3", [][]int{}, 1, []int{0}},
		{"example 4", [][]int{}, 2, []int{0, 1}},
		{"example 5", [][]int{{0, 1}}, 2, []int{1, 0}},
		{"example 6", [][]int{{0, 1}, {1, 0}}, 2, []int{}},
		{"example 7", [][]int{}, 3, []int{0, 1, 2}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := findOrder(tc.numCourses, tc.prereqs)

			assert.Equal(t, tc.want, got)
		})
	}
}
