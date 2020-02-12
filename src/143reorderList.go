package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	slow, quick := head, head
	for quick.Next != nil && quick.Next.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
	}
	head1 := slow.Next
	slow.Next = nil
	var pre *ListNode

	//reverse right
	for head1 != nil {
		tmp := head1.Next
		head1.Next = pre
		pre = head1
		head1 = tmp
	}

	cur1, cur2 := head, pre
	for cur2 != nil {
		next1, next2 := cur1.Next, cur2.Next
		cur1.Next, cur2.Next = cur2, next1
		cur1, cur2 = next1, next2
	}

}
func reorderList1(head *ListNode) {
	if head == nil {
		return
	}
	queue := []*ListNode{}
	cur := head
	for cur != nil {
		queue = append(queue, cur)
		cur = cur.Next
	}
	newHead := &ListNode{0, nil}
	end := newHead
	left := 0
	right := len(queue) - 1
	for left <= right {
		end.Next = queue[left]
		left++
		end = end.Next
		if left <= right {
			end.Next = queue[right]
			right--
			end = end.Next
		}
	}
	end.Next = nil
	head = newHead.Next
}
func main() {
	a1 := &ListNode{1, nil}
	a2 := &ListNode{2, nil}
	a3 := &ListNode{3, nil}
	a4 := &ListNode{4, nil}
	a5 := &ListNode{5, nil}
	a1.Next = a2
	a2.Next = a3
	a3.Next = a4
	a4.Next = a5
	reorderList(a1)
	fmt.Println()
}
