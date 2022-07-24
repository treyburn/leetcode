package matrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_leftMostColumnWithOne(t *testing.T) {
	type testCase struct {
		name  string
		input [][]int
		want  int
	}

	var tests = []testCase{
		{"simple", [][]int{{0, 0}, {1, 1}}, 0},
		{"simple2", [][]int{{0, 0}, {0, 1}}, 1},
		{"none", [][]int{{0, 0}, {0, 0}}, -1},
		{"uneven", [][]int{
			{0, 0, 0, 0, 0},
			{0, 0, 1, 1, 1},
			{0, 0, 0, 1, 1},
			{0, 1, 1, 1, 1},
			{0, 0, 1, 1, 1},
			{0, 0, 0, 0, 1},
		}, 1},
		{"large", [][]int{
			{1, 1, 1, 1, 1},
			{0, 0, 0, 1, 1},
			{0, 0, 1, 1, 1},
			{0, 0, 0, 0, 1},
			{0, 0, 0, 0, 0},
		}, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tc := test

			bm := makeBinaryMatrix(tc.input)

			got := leftMostColumnWithOne(bm)

			assert.Equal(t, tc.want, got)
		})
	}
}

func makeBinaryMatrix(input [][]int) BinaryMatrix {
	return BinaryMatrix{
		Get: func(x int, y int) int {
			return input[x][y]
		},
		Dimensions: func() []int {
			return []int{len(input), len(input[0])}
		},
	}
}

var bigInput = [][]int{
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1},
	{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
}
var bigMatrix = makeBinaryMatrix(bigInput)
var got int

func Benchmark_BinarySearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		got = leftMostColumnWithOne(bigMatrix)
	}
}

// interestingly - the naive iter to scan left is still faster
func Benchmark_Iter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		got = leftMostColumnWithOne_Iter(bigMatrix)
	}
}
