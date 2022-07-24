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
		if col <= 0 {
			break
		}
		left := 0
		got := binaryMatrix.Get(i, col)
		if got == 1 {
			// binary search
			for left <= col {
				mid := (left + col) / 2
				got = binaryMatrix.Get(i, mid)
				if got == 1 {
					col = mid - 1
				} else {
					left = mid + 1
				}
			}
			leftMostCol = max(left, col)
		} else {
			continue
		}
	}

	return leftMostCol
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func leftMostColumnWithOne_Iter(binaryMatrix BinaryMatrix) int {
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
