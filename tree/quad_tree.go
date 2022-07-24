package tree

/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	v, ok := isLeaf(grid)
	if ok {
		return &Node{Val: v, IsLeaf: ok}
	}

	splitLen := len(grid) / 2
	root := &Node{Val: true, IsLeaf: false}
	upperLeft := make([][]int, 0)
	upperRight := make([][]int, 0)
	lowerLeft := make([][]int, 0)
	lowerRight := make([][]int, 0)

	// this parses our inputs into the respective corners
	for rowIdx, subGrid := range grid {
		if rowIdx < splitLen { // write to uppers
			upperLeft = append(upperLeft, subGrid[:splitLen])
			upperRight = append(upperRight, subGrid[splitLen:])
		} else { // write to lowers
			lowerLeft = append(lowerLeft, subGrid[:splitLen])
			lowerRight = append(lowerRight, subGrid[splitLen:])
		}
	}

	root.TopLeft = construct(upperLeft)
	root.TopRight = construct(upperRight)
	root.BottomLeft = construct(lowerLeft)
	root.BottomRight = construct(lowerRight)

	return root
}

func isLeaf(grid [][]int) (bool, bool) {
	value := -1
	valueConv := map[int]bool{0: false, 1: true}
	for _, subgrid := range grid {
		for _, v := range subgrid {
			if value == -1 { // first iter
				value = v
			}
			if v != value {
				return valueConv[value], false
			}
		}
	}

	return valueConv[value], true
}
