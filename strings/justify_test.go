package strings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fullJustify(t *testing.T) {
	type testCase struct {
		name  string
		input []string
		width int
		want  []string
	}

	var tests = []testCase{
		{"example 1", []string{"This", "is", "an", "example", "of", "text", "justification."}, 16, []string{
			"This    is    an",
			"example  of text",
			"justification.  ",
		}},
		{"example 2", []string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16, []string{
			"What   must   be",
			"acknowledgment  ",
			"shall be        ",
		}},
		{"example 3", []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}, 20, []string{
			"Science  is  what we",
			"understand      well",
			"enough to explain to",
			"a  computer.  Art is",
			"everything  else  we",
			"do                  ",
		}},
	}
	
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			got := fullJustify(tc.input, tc.width)

			assert.Equal(t, tc.want, got)
		})
	}
}
