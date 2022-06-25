package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isSubsequence(t *testing.T) {
	type testCase struct {
		name string
		s    string
		t    string
		want bool
	}

	var tests = []testCase{
		{"example 1", "abc", "ahbgdc", true},
		{"example 2", "axc", "ahbgdc", false},
		{"example 3", "ace", "abcde", true},
		{"example 4", "aec", "abcde", false},
		{"example 5", "", "ahbgdc", true},
		{"example 6", "ahbgdc", "", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := isSubsequence(tc.s, tc.t)

			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_isSubsequence2(t *testing.T) {
	type testCase struct {
		name string
		s    string
		t    string
		want bool
	}

	var tests = []testCase{
		{"example 1", "abc", "ahbgdc", true},
		{"example 2", "axc", "ahbgdc", false},
		{"example 3", "ace", "abcde", true},
		{"example 4", "aec", "abcde", false},
		{"example 5", "", "ahbgdc", true},
		{"example 6", "ahbgdc", "", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := isSubsequence2(tc.s, tc.t)

			assert.Equal(t, tc.want, got)
		})
	}
}

var resSub bool
var s1Sub = "abcde"
var t1Sub = "abwqueiqwueoqwieuwqioeuqioweuqiowcjkjwekwqjelkwqdncxmvndfjkvmdcnek"
var s2Sub = "axcde"
var t2Sub = "abwqueiqwueoqwieuwqioeuqioweuqiowcjkjwekwqjelkwqdncqmvndfjkvmdcnek"

func Benchmark_positive_original(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resSub = isSubsequence(s1Sub, t1Sub)
	}
}

func Benchmark_positive_second(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resSub = isSubsequence2(s1Sub, t1Sub)
	}
}

func Benchmark_negative_original(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resSub = isSubsequence(s2Sub, t2Sub)
	}
}

func Benchmark_negative_second(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resSub = isSubsequence2(s2Sub, t2Sub)
	}
}
