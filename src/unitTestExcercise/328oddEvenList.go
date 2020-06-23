package unitTestExcercise

func oddEvenList(head *ListNode) *ListNode {
	one := &ListNode{}
	two := &ListNode{}
	tmp := two
	newHead := one
	i := 1
	for head != nil {
		if i&1 == 1 {
			one.Next = head
			one = one.Next
			head = head.Next
			one.Next = nil
		} else {
			two.Next = head
			two = two.Next
			head = head.Next
			two.Next = nil
		}
		i++
	}
	one.Next = tmp.Next
	return newHead.Next
}
