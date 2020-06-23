package unitTestExcercise

func deleteDuplicates(head *ListNode) *ListNode {
	newHead := &ListNode{}
	newHead.Next = head
	end := newHead
	cur := head
	for cur != nil {
		if cur.Next != nil && cur.Val == cur.Next.Val {
			tmp := cur.Val
			for cur != nil && cur.Val == tmp {
				cur = cur.Next
			}
		} else {
			end.Next = cur
			end = end.Next
			cur = cur.Next
		}
	}
	end.Next = nil
	return newHead.Next
}
