package unitTestExcercise

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	a, b := []*TreeNode{}, []*TreeNode{}
	findNode236(root, p, &a)
	findNode236(root, q, &b)
	return findCommon(a, b)

}

func findNode236(root *TreeNode, target *TreeNode, re *[]*TreeNode) {
	if root == nil {
		return
	}
	if root.Val == target.Val {
		*re = append(*re, root)
		return
	}
	*re = append(*re, root)
	findNode236(root.Left, target, re)
	if (*re)[len(*re)-1].Val == target.Val {
		return
	}
	findNode236(root.Right, target, re)
	if (*re)[len(*re)-1].Val == target.Val {
		return
	}
	*re = (*re)[:len(*re)-1]
}

func findCommon(a, b []*TreeNode) *TreeNode {
	p, q := len(a)-1, len(b)-1
	if len(a)-len(b) > 0 {
		k := len(a) - len(b)
		for k != 0 {
			k--
			p--
		}
	} else {
		k := len(b) - len(a)
		for k != 0 {
			k--
			q--
		}
	}
	for a[p] != b[q] {
		p--
		q--
	}
	return a[p]
}
