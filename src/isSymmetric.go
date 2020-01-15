package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricCore(root.Left, root.Right)
}
func isSymmetricCore(root1 *TreeNode, root2 *TreeNode) bool {
	if root2 == nil && root1 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return isSymmetricCore(root1.Left, root2.Right) && isSymmetricCore(root1.Right, root2.Left)
}
