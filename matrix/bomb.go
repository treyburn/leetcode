package matrix

var (
	W = []byte("W")[0]
	O = []byte("0")[0]
	E = []byte("E")[0]
)

func maxKilledEnemies(grid [][]byte) int {
	count := 0
	for rowIdx, row := range grid {
		for colIdx, tile := range row {
			if tile == O {
				count = _max(count, countNeighbors(rowIdx, colIdx, grid))
			}
		}
	}

	return count
}

func countNeighbors(row, col int, grid [][]byte) int {
	count := 0
	maxX := len(grid[0])
	maxY := len(grid)

	leftX := col - 1
	rightX := col + 1
	upY := row - 1
	downY := row + 1

	for leftX >= 0 || rightX < maxX {
		if leftX >= 0 {
			switch grid[row][leftX] {
			case E:
				count++
				leftX--
			case W:
				leftX = -1
			default:
				leftX--
			}
		}
		if rightX < maxX {
			switch grid[row][rightX] {
			case E:
				count++
				rightX++
			case W:
				rightX = maxX
			default:
				rightX++
			}
		}
	}

	for upY >= 0 || downY < maxY {
		if upY >= 0 {
			switch grid[upY][col] {
			case E:
				count++
				upY--
			case W:
				upY = -1
			default:
				upY--
			}
		}
		if downY < maxY {
			switch grid[downY][col] {
			case E:
				count++
				downY++
			case W:
				downY = maxY
			default:
				downY++
			}
		}
	}

	return count
}

func _max(x, y int) int {
	if x < y {
		return y
	}
	return x
}