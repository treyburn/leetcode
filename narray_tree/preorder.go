package narray_tree

/**
* Definition for a Node.
* type Node struct {
*     Val int
*     Children []*Node
* }
 */

type Node struct {
	Val      int
	Children []*Node
}

// recursive solution
func preorder_recrusive(root *Node) []int {
	if root == nil {
		return nil
	}
	output := []int{root.Val}

	for _, child := range root.Children {
		output = append(output, preorder_recrusive(child)...)
	}

	return output
}

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	output := []int{root.Val}
	kiddos := root.Children

	for len(kiddos) > 0 {
		output = append(output, kiddos[0].Val)
		kiddos = append(kiddos[0].Children, kiddos[1:]...)
	}

	return output
}
