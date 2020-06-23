package unitTestExcercise

func kthSmallest(root *TreeNode, k int) int {
	count := helperkthSmallest(root.Left)
	if count == k-1 {
		return root.Val
	} else if count > k-1 {
		return kthSmallest(root.Left, k)
	} else {
		return kthSmallest(root.Right, k-count-1)
	}
}

func helperkthSmallest(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return helperkthSmallest(root.Left) + helperkthSmallest(root.Right) + 1
}
