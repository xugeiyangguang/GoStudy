package main

func swapPairs(head *ListNode) *ListNode {
	newHead := &ListNode{0, nil}
	newHead.Next = head
	cur := newHead
	for cur.Next != nil && cur.Next.Next != nil {
		next1 := cur.Next
		next2 := cur.Next.Next
		next3 := cur.Next.Next.Next
		cur.Next = next2
		next2.Next = next1
		next1.Next = next3
		cur = next1
	}
	return newHead.Next
}
