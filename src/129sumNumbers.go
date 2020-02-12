package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

func main() {

	a2 := &TreeNode{2, nil, nil}
	a3 := &TreeNode{3, nil, nil}
	a1 := &TreeNode{1, a2, a3}
	fmt.Println(sumNumbers(a1))
}
