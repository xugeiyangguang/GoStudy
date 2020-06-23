package unitTestExcercise

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	newHead := new(ListNode)
	end := newHead
	p1 := l1
	p2 := l2
	last := 0 //表示上次需要进位的值
	for p1 != nil && p2 != nil {
		val := p1.Val + p2.Val + last
		if val >= 10 {
			last = 1
		} else {
			last = 0
		}
		tmp := &ListNode{val % 10, nil}
		end.Next = tmp
		end = end.Next
		p1 = p1.Next
		p2 = p2.Next
	}
	var remain *ListNode
	if p1 != nil {
		remain = p1
	} else {
		remain = p2
	}
	for remain != nil {
		val := remain.Val + last
		if val >= 10 {
			last = 1
		} else {
			last = 0
		}
		tmp := &ListNode{val % 10, nil}
		end.Next = tmp
		end = end.Next
		remain = remain.Next
	}
	if last == 1 {
		tmp := &ListNode{1, nil}
		end.Next = tmp
	}
	return newHead.Next
}
