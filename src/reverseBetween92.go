package main

import (
	"fmt"
	"node"
)

func reverseBetween(head *node.ListNode, m int, n int) *node.ListNode {
	newHead := &node.ListNode{0, nil}
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
	a1 := &node.ListNode{1, nil}
	a2 := &node.ListNode{2, nil}
	a3 := &node.ListNode{3, nil}
	a4 := &node.ListNode{4, nil}
	a5 := &node.ListNode{5, nil}
	a1.Next = a2
	a2.Next = a3
	a3.Next = a4
	a4.Next = a5
	re := reverseBetween(a1, 1, 4)
	fmt.Println(re)

}
