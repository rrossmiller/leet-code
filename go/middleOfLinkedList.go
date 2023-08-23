package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	x := []int{6, 5, 4, 3, 2, 1}
	var head *ListNode
	for _, i := range x {
		head = &ListNode{i, head}
	}

	h := middleNode(head)
	fmt.Println(h, "should be 4")

	fmt.Println()
	x = []int{5, 4, 3, 2, 1}
	head = nil
	for _, i := range x {
		head = &ListNode{i, head}
	}

	h = middleNode(head)
	fmt.Println(h, "should be 3")

	fmt.Println()
	x = []int{3, 2, 1}
	head = nil
	for _, i := range x {
		head = &ListNode{i, head}
	}

	h = middleNode(head)
	fmt.Println(h, "should be 2")
}

func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow 
}

// func middleNode(head *ListNode) *ListNode {
// 	l := 0
// 	for hld := head; hld != nil; {
// 		l++
// 		hld = hld.Next
//
// 	}
//
// 	fmt.Println(l/2, l)
// 	for x := l / 2; x > 0; x-- {
// 		fmt.Println(x)
// 		head = head.Next
// 	}
//
// 	return head
// }
