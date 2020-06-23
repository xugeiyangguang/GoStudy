package unitTestExcercise

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	one := helper337(root.Left, 0) + helper337(root.Right, 0) + root.Val
	two := helper337(root.Left, 1) + helper337(root.Right, 1)
	re := Max(one, two)
	return re
}

func helper337(root *TreeNode, can int) int {
	if root == nil {
		return 0
	}
	cur := 0
	if can == 1 {
		one := helper337(root.Left, 0) + helper337(root.Right, 0) + root.Val
		two := helper337(root.Left, 1) + helper337(root.Right, 1)
		cur = Max(one, two)
	} else {
		cur = helper337(root.Left, 1) + helper337(root.Right, 1)
	}
	return cur
}
