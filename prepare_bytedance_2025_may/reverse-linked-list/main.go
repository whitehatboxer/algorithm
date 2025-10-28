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

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode = nil
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

func reverseListOld(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode
	var tmp *ListNode

	current := head
	for {
		tmp = current.Next
		current.Next = prev
		if tmp == nil {
			break
		} else {
			prev = current
			current = tmp
		}
	}

	return current
}

func reverseListNew(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode
	current := head

	for {
		next := current.Next
		current.Next = prev
		prev = current

		if next == nil {
			break
		}
		current = next
	}

	return current
}

func main() {
	ln := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	ln.show()
	rln := reverseListNew(ln)
	rln.show()

	fmt.Println("====")
	ln1 := &ListNode{
		Val:  5,
		Next: nil,
	}
	ln1.show()
	rln1 := reverseListNew(ln1)
	rln1.show()
}
