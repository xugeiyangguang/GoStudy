package unitTestExcercise

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	re := 0
	helper2(root, 0, &re)
	return re
}

func helper2(root *TreeNode, tmp int, re *int) {
	if root == nil {
		return
	}
	tmp = 10*tmp + root.Val
	if root.Left == nil && root.Right == nil {
		*re += tmp
		return
	}
	helper2(root.Left, tmp, re)
	helper2(root.Right, tmp, re)
}
