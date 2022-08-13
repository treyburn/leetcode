package destination

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_destCity(t *testing.T) {
	type testCase struct {
		name  string
		input [][]string
		want  string
	}

	var tests = []testCase{
		{"example 1", [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}, "Sao Paulo"},
		{"example 2", [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}, "A"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := destCity(tc.input)

			assert.Equal(t, tc.want, got)
		})
	}
}
