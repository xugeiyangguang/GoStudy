package unitTestExcercise

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	one, two := head, head
	for {
		if two == nil || two.Next == nil {
			return nil
		}
		one = one.Next
		two = two.Next.Next
		if one == two {
			break
		}
	}
	for head != two {
		head = head.Next
		two = two.Next
	}
	return head
}
