package unitTestExcercise

import "sort"

func flatten(root *TreeNode) *TreeNode {
	nodes := []*TreeNode{}
	findNode(root, &nodes)

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Val < nodes[j].Val
	})

	head := &TreeNode{0, nil, nil}
	last := head
	head.Left = nil
	for i := 0; i < len(nodes); i++ {
		last.Right = nodes[i]
		nodes[i].Left = nil
		last = last.Right
	}
	last.Right = nil
	last.Left = nil
	root = head.Right
	return root
}

func findNode(root *TreeNode, nodes *[]*TreeNode) {
	if root != nil {
		*nodes = append(*nodes, root)
		findNode(root.Left, nodes)
		findNode(root.Right, nodes)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createNode() *TreeNode {
	/*	a3 := &TreeNode{3, nil, nil}
		a4 := &TreeNode{4, nil, nil}
		a6 := &TreeNode{6, nil, nil}
		a2 := &TreeNode{2, a3, a4}
		a5 := &TreeNode{5, nil, a6}
		a1 := &TreeNode{1, a2, a5}*/
	a5 := &TreeNode{5, nil, nil}
	a3 := &TreeNode{3, a5, nil}
	a4 := &TreeNode{4, nil, nil}
	a2 := &TreeNode{2, a3, a4}
	a1 := &TreeNode{1, a2, nil}
	return a1
}
