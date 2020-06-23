package unitTestExcercise

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	re := [][]int{}
	cur := []*TreeNode{root}
	value := []int{root.Val}
	re = append(re, value)
	i := 1
	for len(cur) != 0 {
		i = -i
		value := []int{}
		next := make([]*TreeNode, 0)
		for i := range cur {
			if cur[i].Left != nil {
				next = append(next, cur[i].Left)
				value = append(value, cur[i].Left.Val)
			}
			if cur[i].Right != nil {
				next = append(next, cur[i].Right)
				value = append(value, cur[i].Right.Val)
			}
		}
		cur = next
		if len(value) != 0 {
			if i == -1 {
				l := len(value)
				for k := 0; k < l/2; k++ {
					value[k], value[l-k-1] = value[l-k-1], value[k]
				}
			}
			re = append(re, value)
		}

	}
	return re
}
