package main

import "fmt"

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func (l *LinkNode) Show() {
	for l != nil {
		fmt.Print(l.Val, ">")
		l = l.Next
	}
	fmt.Println()
}

func reverse(head *LinkNode) *LinkNode {
	if head == nil {
		return nil
	}

	current := head
	var prev *LinkNode

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

func reverse1(head *LinkNode) *LinkNode {
	if head == nil {
		return nil
	}

	current := head
	var prev *LinkNode
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

// add in 2025.5.11
func reverse2(head *LinkNode) *LinkNode {
	if head == nil {
		return nil
	}

	current := head
	var prev *LinkNode
	for {
		next := current.Next
		current.Next = prev
		prev = current
		if next == nil {
			break
		} else {
			current = next
		}
	}

	return current
}

func main() {
	node := &LinkNode{
		Val: 1,
		Next: &LinkNode{
			Val: 2,
			Next: &LinkNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	newNode := reverse2(node)
	newNode.Show()
}
