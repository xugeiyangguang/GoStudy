package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	end := &ListNode{0, nil}
	start := &ListNode{0, nil}
	end.Next = start
	pos := head
	for pos != nil {
		tmp := pos.Next
		cur := end
		for {
			if cur.Next == start || cur.Next.Val < pos.Val {
				pos.Next = cur.Next
				cur.Next = pos
				break
			} else {
				cur = cur.Next
			}
		}
		pos = tmp
	}
	start = nil
	re := reverseList(end)
	return re.Next
}

func reverseList(head *ListNode) *ListNode {
	newHead := &ListNode{}
	newHead.Next = nil
	cur := head.Next
	for cur != nil {
		tmp := cur.Next
		cur.Next = newHead.Next
		newHead.Next = cur
		cur = tmp
	}
	return newHead.Next
}

func main() {
	a1 := &ListNode{-1, nil}
	a2 := &ListNode{5, nil}
	a3 := &ListNode{3, nil}
	a4 := &ListNode{4, nil}
	a5 := &ListNode{0, nil}
	a1.Next = a2
	a2.Next = a3
	a3.Next = a4
	a4.Next = a5
	re := insertionSortList(a1)
	fmt.Println(re)
}
