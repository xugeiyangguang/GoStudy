package unitTestExcercise

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	start := &ListNode{}
	start.Next = head
	slow := start
	quick := start
	for i := 0; i < n; i++ {
		quick = quick.Next
	}
	for quick.Next != nil {
		slow = slow.Next
		quick = quick.Next
	}
	slow.Next = slow.Next.Next
	return start.Next
}
