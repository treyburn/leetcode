package logger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogger_ShouldPrintMessage(t *testing.T) {
	type input struct {
		timestamp int
		message   string
	}

	type testCase struct {
		name  string
		input []input
		want  []bool
	}

	var tests = []testCase{
		{"example 1", []input{
			{1, "foo"},
			{2, "bar"},
			{3, "foo"},
			{8, "bar"},
			{10, "foo"},
			{11, "foo"},
		}, []bool{true, true, false, false, false, true}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			l := Constructor()

			for idx, msg := range tc.input {
				got := l.ShouldPrintMessage(msg.timestamp, msg.message)

				assert.Equal(t, tc.want[idx], got)
			}
		})
	}
}
