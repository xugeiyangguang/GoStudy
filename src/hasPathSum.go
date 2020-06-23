package src

import "unitTestExcercise"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func hasPathSum(root *unitTestExcercise.TreeNode, sum int) bool {
	if sum < 0 || root == nil {
		return false
	}
	if sum == 0 && root.Left == nil && root.Right == nil {
		return true
	}

	return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
}
