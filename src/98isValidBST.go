package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

func main() {
	a1 := &TreeNode{3, nil, nil}
	a2 := &TreeNode{1, nil, nil}
	a3 := &TreeNode{2, nil, nil}
	a4 := &TreeNode{3, nil, nil}
	a1.Left = a2
	a2.Right = a3
	a3.Right = a4
	re := isValidBST(a1)
	fmt.Println(re)
}
