package test

import "node"

func mergeTwoLists(l1 *node.ListNode, l2 *node.ListNode) *node.ListNode {
	head := &node.ListNode{0, nil}
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

}
