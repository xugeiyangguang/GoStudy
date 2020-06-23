package src

import (
	"fmt"
	"unitTestExcercise"
)

func reverseBetween(head *unitTestExcercise.node.ListNode, m int, n int) *unitTestExcercise.node.ListNode {
	newHead := &unitTestExcercise.node.ListNode{0, nil}
	newHead.Next = head
	end := newHead
	i := 0
	for ; i < m-1; i++ {
		end = end.Next
	}
	cur := end.Next
	i++
	skip := cur
	end.Next = nil
	for ; i <= n; i++ {
		tmp := cur
		cur = cur.Next
		tmp.Next = end.Next
		end.Next = tmp
		if i == n {
			skip.Next = cur
			return newHead.Next
		}
	}
	return newHead.Next
}

func main() {
	a1 := &unitTestExcercise.node.ListNode{1, nil}
	a2 := &unitTestExcercise.node.ListNode{2, nil}
	a3 := &unitTestExcercise.node.ListNode{3, nil}
	a4 := &unitTestExcercise.node.ListNode{4, nil}
	a5 := &unitTestExcercise.node.ListNode{5, nil}
	a1.Next = a2
	a2.Next = a3
	a3.Next = a4
	a4.Next = a5
	re := reverseBetween(a1, 1, 4)
	fmt.Println(re)

}
