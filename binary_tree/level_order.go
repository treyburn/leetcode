package binary_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var output = make([][]int, 0)
	var nextLevel = make([]*TreeNode, 0)
	output = append(output, []int{root.Val})
	if root.Left != nil {
		nextLevel = append(nextLevel, root.Left)
	}
	if root.Right != nil {
		nextLevel = append(nextLevel, root.Right)
	}
	output = append(output, traverseLevel(nextLevel)...)
	return output
}

func traverseLevel(level []*TreeNode) [][]int {
	if len(level) == 0 {
		return nil
	}
	var finalOutput = make([][]int, 0)
	var thisLevelsOutput = make([]int, 0)
	var nextLevel = make([]*TreeNode, 0)

	for _, branch := range level {
		thisLevelsOutput = append(thisLevelsOutput, branch.Val)
		if branch.Left != nil {
			nextLevel = append(nextLevel, branch.Left)
		}
		if branch.Right != nil {
			nextLevel = append(nextLevel, branch.Right)
		}
	}

	finalOutput = append(finalOutput, thisLevelsOutput)
	finalOutput = append(finalOutput, traverseLevel(nextLevel)...)
	return finalOutput
}
