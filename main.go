package main

import "fmt"

type entry struct {
	val, min int
}

// MinStack is a stack that also answers "what is the minimum right now?"
// in O(1): each entry carries the minimum of the stack up to itself, so
// popping needs no recomputation.
type MinStack struct {
	data []entry
}

// Push adds v, remembering the running minimum alongside it.
func (s *MinStack) Push(v int) {
	m := v
	if len(s.data) > 0 && s.data[len(s.data)-1].min < v {
		m = s.data[len(s.data)-1].min
	}
	s.data = append(s.data, entry{v, m})
}

// Pop removes and returns the top value.
func (s *MinStack) Pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	v := s.data[len(s.data)-1].val
	s.data = s.data[:len(s.data)-1]
	return v, true
}

// Top returns the top value without removing it.
func (s *MinStack) Top() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	return s.data[len(s.data)-1].val, true
}

// Min returns the smallest value currently on the stack.
func (s *MinStack) Min() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	return s.data[len(s.data)-1].min, true
}

func main() {
	var s MinStack
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	fmt.Println(s.Min()) // -3
	s.Pop()
	fmt.Println(s.Top()) // 0
	fmt.Println(s.Min()) // -2
}
