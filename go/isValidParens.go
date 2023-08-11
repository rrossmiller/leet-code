package main

import "fmt"

func main() {
	fmt.Println(isValid("([{}])"))
	fmt.Println(isValid("()"))
}

type (
	Stack[T any] struct {
		top *node[T]
		len int
	}
	node[T any] struct {
		value T
		next  *node[T]
	}
)

func (s *Stack[T]) Len() int {
	return s.len
}
func (s *Stack[T]) Peek() T {
	if s.len == 0 {
		var t T
		return t
	}
	return s.top.value
}
func (s *Stack[T]) Pop() T {
	if s.len == 0 {
		var t T
		return t
	}
	n := s.top
	s.top = n.next
	s.len--
	return n.value
}
func (s *Stack[T]) Push(val T) {
	n := &node[T]{val, s.top}
	s.top = n
	s.len++
}

func newRuneStack() *Stack[rune] {
	return &Stack[rune]{nil, 0}
}

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}

	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	stack := newRuneStack()

	for _, v := range s {
		// If we see an opening parenthesis we put it to stack
		if _, ok := pairs[v]; ok {
			stack.Push(v)
			// If it’s the closing one, then it must correspond to the parenthesis at the top of our stack, so we check if it is true and we remove this pair of parentheses if it's so. If not, we return false and stop the function’s execution immediately.
		} else if pairs[stack.Peek()] != v || stack.len == 0 {
			return false
		} else {
			stack.Pop()
		}

	}
	return stack.len == 0
}
