package unitTestExcercise

import (
	"math"
)

func isValidBST(root *TreeNode) bool {
	return helper6(root, math.MinInt64, math.MaxInt64)
}

func helper6(root *TreeNode, left, right int) bool {
	if root == nil {
		return true
	}
	if root.Val < left || root.Val > right {
		return false
	}
	if root.Left != nil && root.Left.Val >= root.Val || root.Right != nil && root.Right.Val <= root.Val {
		return false
	}
	return helper6(root.Left, left, root.Val) && helper6(root.Right, root.Val, right)
}
