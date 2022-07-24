package matrix

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxKilledEnemies(t *testing.T) {
	type testCase struct {
		name  string
		input [][]string
		want  int
	}

	var tests = []testCase{
		{"simple",
			[][]string{
				{"0", "E", "0", "0"},
				{"E", "0", "W", "E"},
				{"0", "E", "0", "0"},
			}, 3},
		{"only 1",
			[][]string{
				{"W", "W", "W"},
				{"0", "0", "0"},
				{"E", "E", "E"},
			}, 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			grid := buildGrid(tc.input)

			got := maxKilledEnemies(grid)

			assert.Equal(t, tc.want, got)
		})
	}

}

func buildGrid(input [][]string) [][]byte {
	output := make([][]byte, 0)

	for _, row := range input {
		subgrid := make([]byte, 0)
		for _, letter := range row {
			b := []byte(letter)
			if len(b) > 1 {
				panic(fmt.Sprintf("bad conversion for %v", letter))
			}
			subgrid = append(subgrid, b[0])
		}
		output = append(output, subgrid)
	}
	return output
}
