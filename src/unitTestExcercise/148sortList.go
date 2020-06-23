package unitTestExcercise

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	one, two := head, head
	for two.Next != nil && two.Next.Next != nil {
		two = two.Next.Next
		one = one.Next
	}
	newHead := one.Next
	one.Next = nil
	a := sortList(head)
	b := sortList(newHead)
	re := &ListNode{}
	cur := re
	for a != nil && b != nil {
		if a.Val < b.Val {
			cur.Next = a
			a = a.Next
		} else {
			cur.Next = b
			b = b.Next
		}
		cur = cur.Next
	}
	if a == nil {
		cur.Next = b
	}
	if b == nil {
		cur.Next = a
	}
	return re.Next
}
