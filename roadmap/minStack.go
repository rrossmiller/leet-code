package main

import (
	"math"
)

type MinStack struct {
	stack []int
	min   int
}

func Constructor() MinStack {
	return MinStack{[]int{}, math.MaxInt}
}

func (s *MinStack) Push(val int) {
	s.stack = append(s.stack, val)
	if val < s.min {
		s.min = val
	}
}

func (s *MinStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
	s.min = math.MaxInt
	for _, v := range s.stack {
		if v < s.min {
			s.min = v
		}
	}
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	return s.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
