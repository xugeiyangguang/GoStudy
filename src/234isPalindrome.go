package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*func isPalindrome(head *ListNode) bool {
	if head==nil{
		return false
	}
	one := head
	two := head
	slice := make([]int, 0)
	for two.Next != nil && two.Next.Next != nil {
		slice = append(slice, one.Val)
		one = one.Next
		two = two.Next.Next
	}
	if two != nil && two.Next == nil {
	} else {
		slice = append(slice, one.Val)
	}
	one = one.Next
	i := len(slice) - 1
	for ; one != nil && i >= 0; i-- {
		if slice[i] != one.Val {
			return false
		}
		one = one.Next
	}
	if i < 0 && one == nil {
		return true
	}
	return false
}*/

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	one := head
	two := head
	newHead := &ListNode{0, nil}
	var tmp *ListNode
	for two != nil && two.Next != nil {
		two = two.Next.Next
		tmp = one.Next
		one.Next = newHead.Next
		newHead.Next = one
		one = tmp
	}

	if two != nil && two.Next == nil {
		one = one.Next
	}

	newHead = newHead.Next
	for newHead != nil && one != nil {
		if newHead.Val != one.Val {
			return false
		}
		newHead = newHead.Next
		one = one.Next
	}
	if one == nil && newHead == nil {
		return true
	}
	return false
}

func main() {
	a1 := &ListNode{1, nil}
	a2 := &ListNode{2, nil}
	a3 := &ListNode{1, nil}
	//a4 := &ListNode{1, nil}
	//	a5 := &ListNode{1, nil}
	a1.Next = a2
	a2.Next = a3
	//a3.Next = a4
	//	a4.Next = a5
	fmt.Println(isPalindrome(a1))
}
