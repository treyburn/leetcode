package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	type testCase struct {
		name  string
		nums1 []int
		nums2 []int
		want  float64
	}

	var tests = []testCase{
		{"1", []int{1, 3}, []int{2}, 2.0},
		{"2", []int{1, 2}, []int{3, 4}, 2.5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test
			got := findMedianSortedArrays(tc.nums1, tc.nums2)
			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_merge(t *testing.T) {
	type testCase struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}

	var tests = []testCase{
		{"even", []int{1, 3}, []int{2, 4}, []int{1, 2, 3, 4}},
		{"odd", []int{1, 3}, []int{2}, []int{1, 2, 3}},
		{"first empty", []int{}, []int{2}, []int{2}},
		{"second empty", []int{2}, []int{}, []int{2}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test
			got := merge(tc.nums1, tc.nums2)
			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_median(t *testing.T) {
	type testCase struct {
		name string
		nums []int
		want float64
	}

	var tests = []testCase{
		{"even", []int{1, 2, 3, 4}, 2.5},
		{"odd", []int{1, 2, 3}, 2.0},
		{"one element", []int{2}, 2.0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test
			got := median(tc.nums)
			assert.Equal(t, tc.want, got)
		})
	}
}
