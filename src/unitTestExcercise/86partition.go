package unitTestExcercise

func partition(head *ListNode, x int) *ListNode {
	less, more := &ListNode{}, &ListNode{}
	p, q := less, more
	for head != nil {

		if head.Val < x {
			p.Next = head
			p = p.Next
		} else {
			q.Next = head
			q = q.Next
		}
		head = head.Next
	}
	p.Next = more.Next
	q.Next = nil
	return less.Next
}
