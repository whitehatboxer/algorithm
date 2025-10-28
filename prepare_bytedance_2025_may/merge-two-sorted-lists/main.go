package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) show() {
	for l != nil {
		fmt.Print(l.Val, ">")
		l = l.Next
	}
	fmt.Println()
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	s1 := list1
	s2 := list2

	start := &ListNode{
		Next: nil,
	}
	current := start
	for s1 != nil && s2 != nil {
		if s1.Val > s2.Val {
			current.Next = s2
			s2 = s2.Next
		} else {
			current.Next = s1
			s1 = s1.Next
		}
		current = current.Next
	}
	if s1 == nil {
		current.Next = s2
	} else {
		current.Next = s1
	}
	return start.Next

	// var current *ListNode
	// if s1.Val >= s2.Val {
	// 	current = s2
	// 	s2 = s2.Next
	// } else {
	// 	current = s1
	// 	s1 = s1.Next
	// }
	// start := current
	// for s1 != nil && s2 != nil {
	// 	if s1.Val >= s2.Val {
	// 		current.Next = s2
	// 		s2 = s2.Next
	// 	} else {
	// 		current.Next = s1
	// 		s1 = s1.Next
	// 	}
	// 	current = current.Next
	// }
	// if s1 != nil {
	// 	current.Next = s1
	// } else {
	// 	current.Next = s2
	// }
	// return start
}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}

	l1.show()
	l2.show()
	ll := mergeTwoLists(l1, l2)
	ll.show()
}
