package main

import (
	"fmt"
)

func main() {
	head := &ListNode{5, nil}
	for i := 4; i >= 0; i-- {
		head = &ListNode{i, head}
	}

	for t := head; t != nil; {
		fmt.Println(t)
		t = t.Next
	}
	fmt.Println()
	for n := reverseList(head); n != nil; {
		fmt.Println(n, n.Next)
		n = n.Next
	}

	fmt.Println()
	fmt.Println()
	head = &ListNode{2, nil}
	head = &ListNode{1, head}

	for t := head; t != nil; {
		fmt.Println(t)
		t = t.Next
	}
	fmt.Println()
	for n := reverseList(head); n.Next != nil; {
		fmt.Println(n, n.Next)
		n = n.Next
	}
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}
