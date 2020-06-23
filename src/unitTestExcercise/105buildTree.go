package unitTestExcercise

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	a := &TreeNode{Val: preorder[0]}
	i := find_index(inorder, preorder[0])
	a.Left = buildTree(preorder[1:i+1], inorder[:i])
	a.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return a
}
func find_index(a []int, b int) int {
	for i := range a {
		if a[i] == b {
			return i
		}
	}
	return -1
}
