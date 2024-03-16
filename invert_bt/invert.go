package invert

/*
	Given the root of a binary tree, invert the tree, and return its root.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	var walk func(r *TreeNode)
	walk = func(r *TreeNode) {
		if r == nil {
			return
		}
		temp := r.Left
		r.Left = r.Right
		r.Right = temp
		walk(r.Left)
		walk(r.Right)
	}
	walk(root)
	return root
}
