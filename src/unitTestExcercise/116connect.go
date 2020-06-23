package unitTestExcercise

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		root.Left.Next = root.Right
		if root.Next != nil {
			root.Right.Next = root.Next.Left
		}
	}
	connect(root.Left)
	connect(root.Right)
	return root
}
func connect1(root *Node) *Node {
	if root == nil {
		return nil
	}
	stack := make([]*Node, 0)
	stack = append(stack, root)
	curLevel := 1
	nextLevel := 0
	for curLevel != 0 {
		last := stack[0]
		stack = stack[1:]
		if last.Left != nil {
			stack = append(stack, last.Left)
			nextLevel++
		}
		if last.Right != nil {
			stack = append(stack, last.Right)
			nextLevel++
		}
		curLevel--
		if curLevel != 0 {
			last.Next = stack[0]
		} else {
			last.Next = nil
			curLevel, nextLevel = nextLevel, 0
		}
	}
	return root
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
