package unitTestExcercise

func countNodes1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	curLevel := 1
	nextLevel := 0
	re := curLevel
	for len(stack) != 0 {
		curLevel--
		if stack[0].Left != nil {
			stack = append(stack, stack[0].Left)
			nextLevel++
		}
		if stack[0].Right != nil {
			stack = append(stack, stack[0].Right)
			nextLevel++
		}
		stack = stack[1:]
		if curLevel == 0 {
			curLevel = nextLevel
			re += nextLevel
			nextLevel = 0
		}
	}
	return re
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var count int
	fullHeight := findHeight(root)
	for root != nil {
		fullHeight--
		rightHeight := findHeight(root.Right)
		if rightHeight == fullHeight {
			count += 1 << uint(fullHeight)
			root = root.Right
		} else {
			count += 1 << uint(fullHeight-1)
			root = root.Left
		}
	}

	return count
}
func findHeight(root *TreeNode) int {
	var height int
	for root != nil {
		root = root.Left
		height++
	}
	return height
}
