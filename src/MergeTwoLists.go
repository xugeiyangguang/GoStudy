package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	cur := head
	p := l1
	q := l2
	for p != nil && q != nil {
		if p.Val < q.Val {
			cur.Next = p
			p = p.Next
		} else {
			cur.Next = q
			q = q.Next
		}
		cur = cur.Next
	}
	if p == nil {
		cur.Next = q
	} else {
		cur.Next = p
	}
	return head.Next
}

func main() {
	me
}
