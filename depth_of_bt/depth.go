package depth

/*
Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	res := 0

	if root == nil {
		return res
	}

	var walk func(node *TreeNode, counter int)
	walk = func(node *TreeNode, counter int) {
		if node.Left == nil && node.Right == nil {
			if counter > res {
				res = counter
				return
			}
		}

		if node.Left != nil {
			walk(node.Left, counter+1)
		}

		if node.Right != nil {
			walk(node.Right, counter+1)
		}
	}

	walk(root, 1)

	return res
}
