package main

import (
	"fmt"
	"math"
	// "strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// func main() {
// 	// https://leetcode.com/problems/add-two-numbers/
// 	// [2,4,3]
// 	// [5,6,4]
// 	// [7,0,8]
// 	var r1, r2 *ListNode
// 	r1, r2 = nil, nil
// 	l1 := []int{2, 4, 3}
// 	for i := len(l1) - 1; i >= 0; i-- {
// 		r1 = &ListNode{l1[i], r1}
//
// 	}
// 	l2 := []int{5, 6, 4}
//
// 	for i := len(l2) - 1; i >= 0; i-- {
// 		r2 = &ListNode{l2[i], r2}
//
// 	}
// 	ans := addTwoNumbers(r1, r2)
// 	out := "["
// 	for ans != nil {
// 		out += strconv.Itoa(ans.Val)
// 		ans = ans.Next
// 		if ans != nil {
// 			out += ","
// 		}
// 	}
// 	fmt.Println(out + "]")
// 	fmt.Println("[7,0,8]")
// 	fmt.Println("************")
//
// 	// [1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]
// 	// [5,6,4]
// 	r1, r2 = nil, nil
// 	l1 = []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}
// 	for i := len(l1) - 1; i >= 0; i-- {
// 		r1 = &ListNode{l1[i], r1}
//
// 	}

// 	l2 = []int{5, 6, 4}
// 	for i := len(l2) - 1; i >= 0; i-- {
// 		r2 = &ListNode{l2[i], r2}
//
// 	}
//
// 	ans = addTwoNumbers(r1, r2)
// 	out = "["
// 	for ans != nil {
// 		out += strconv.Itoa(ans.Val)
// 		ans = ans.Next
// 		if ans != nil {
// 			out += ","
// 		}
// 	}
// 	fmt.Println(out + "]")
// 	fmt.Println("[6,6,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]")
// 	fmt.Println("************")
//
// 	// l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// 	// Output: [8,9,9,9,0,0,0,1]
//
// }

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var rtn *ListNode
	ans := make([]int, 0)

	carry := 0
	for i := 0; l1 != nil || l2 != nil; i++ {
		elem := 0
		if l1 != nil {
			elem += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			elem += l2.Val
			l2 = l2.Next
		}

		if carry > 0 {
			elem += carry
			carry = 0
		}
		if elem > 9 {
			carry = int(elem / 10)
			fmt.Println(elem, carry)
			elem = elem - carry*10
		}
		ans = append(ans, elem)
	}

	if carry > 0 {
		ans = append(ans, carry)
	}

	for i := len(ans) - 1; i >= 0; i-- {
		rtn = &ListNode{ans[i], rtn}
	}
	return rtn
}
func work2(l1 *ListNode, l2 *ListNode) []int {
	rtn := make([]int, 0)
	carry := 0
	for i := 0; l1 != nil || l2 != nil; i++ {
		elem := 0
		if l1 != nil {
			elem += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			elem += l2.Val
			l2 = l2.Next
		}

		if carry > 0 {
			elem += carry
			carry = 0
		}
		if elem > 9 {
			carry = int(elem / 10)
			fmt.Println(elem, carry)
			elem = elem - carry*10
		}
		rtn = append(rtn, elem)
	}

	if carry > 0 {
		rtn = append(rtn, carry)
	}
	return rtn
}

func work(l *ListNode) int {
	rtn := 0 //make([]int, 0)
	place := 0
	for l != nil {
		pow := math.Pow10(place)
		add := l.Val * int(pow)
		rtn += add
		place++
		l = l.Next
	}
	return rtn
}
