package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var temp *ListNode

	var current *ListNode
	var prev *ListNode

	current = head
	for {
		if current.Next != nil {
			temp = current.Next
			current.Next = prev
			prev = current
			current = temp
		} else {
			current.Next = prev
			break
		}
	}

	return current
}


func main() {
	ln := ListNode{
		Val: 0,
		Next: nil,
	}
	ln1 := ListNode{
		Val: 1,
		Next: &ln,
	}
	ln2 := ListNode{
		Val: 2,
		Next: &ln1,
	}
	ln3 := ListNode{
		Val: 3,
		Next: &ln2,
	}
	ln4 := ListNode{
		Val: 4,
		Next: &ln3,
	}

	// notice ln4 is head
	newHead := reverseList(&ln4)
	fmt.Println(newHead.Val)
	fmt.Println(newHead.Next.Val)
}
