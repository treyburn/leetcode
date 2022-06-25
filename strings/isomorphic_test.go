package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isIsomorphic(t *testing.T) {
	type testCase struct {
		name string
		s1   string
		s2   string
		want bool
	}

	var tests = []testCase{
		{"example 1", "egg", "add", true},
		{"example 2", "foo", "bar", false},
		{"example 3", "paper", "title", true},
		{"example 4", "papap", "titii", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := isIsomorphic(tc.s1, tc.s2)

			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_isIsomorphic2(t *testing.T) {
	type testCase struct {
		name string
		s1   string
		s2   string
		want bool
	}

	var tests = []testCase{
		{"example 1", "egg", "add", true},
		{"example 2", "foo", "bar", false},
		{"example 3", "paper", "title", true},
		{"example 4", "papap", "titii", false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := isIsomorphic2(tc.s1, tc.s2)

			assert.Equal(t, tc.want, got)
		})
	}
}

var res bool
var s1 = "paperpaperpaperpaperpaperpaperpaperpaperpaperpaperpaperpaperpaperpaperpaperpaper"
var t1 = "titletitletitletitletitletitletitletitletitletitletitletitletitletitletitletitle"
var s2 = "paperpaperpaperpaperpaperpaperpaperpaperpaperpaperpaperpaperpaperpaperpaperpaper"
var t2 = "titletitletitletitletitletitletitletitletitletitletitletitletitletitletitletitlq"

func Benchmark_Positive_Original(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res = isIsomorphic(s1, t1)
	}
}

func Benchmark_Positive_Second(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res = isIsomorphic2(s1, t1)
	}
}

func Benchmark_Negative_Original(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res = isIsomorphic(s2, t2)
	}
}

func Benchmark_Negative_Second(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res = isIsomorphic2(s2, t2)
	}
}
