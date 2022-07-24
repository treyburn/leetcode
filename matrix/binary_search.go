package matrix

type BinaryMatrix struct {
	Get        func(int, int) int
	Dimensions func() []int
}

func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
	leftMostCol := -1
	dims := binaryMatrix.Dimensions()
	row := dims[0] - 1
	col := dims[1] - 1
	for i := 0; i <= row; i++ {
		for col >= 0 {
			got := binaryMatrix.Get(i, col)
			if got == 1 {
				leftMostCol = col
				col--
			} else {
				break
			}
		}
	}

	return leftMostCol
}
