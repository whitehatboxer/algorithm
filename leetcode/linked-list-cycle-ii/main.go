/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	distinctMap := make(map[*ListNode]int)

	for {
		if head == nil {
			return nil
		}

		if _, ok := distinctMap[head]; ok {
			return head
		}

		distinctMap[head] = 1
		head = head.Next
	}
}

func main() {
	l1 := &ListNode{
		Val:  1,
		Next: nil,
	}

	l2 := new(ListNode)
	l2.Next = l1
	l2.Val = 12

	l3 := new(ListNode)
	l3.Next = l2
	l3.Val = 11

	l4 := new(ListNode)
	l4.Next = l3
	l4.Val = 3

	l1.Next = l3

	res := detectCycle(l4)
	fmt.Println(res)
}
