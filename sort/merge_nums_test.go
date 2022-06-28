package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_merge(t *testing.T) {
	type testCase struct {
		name  string
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}

	var tests = []testCase{
		{"example 1", []int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{"example 2", []int{1}, 1, []int{}, 0, []int{1}},
		{"example 3", []int{0}, 0, []int{1}, 1, []int{1}},
		{"example 4", []int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3, []int{1, 2, 3, 4, 5, 6}},
		{"example 5", []int{1, 5, 7, 0, 0, 0}, 3, []int{1, 5, 6}, 3, []int{1, 1, 5, 5, 6, 7}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			merge(tc.nums1, tc.m, tc.nums2, tc.n)

			assert.Equal(t, tc.want, tc.nums1)
		})
	}
}
