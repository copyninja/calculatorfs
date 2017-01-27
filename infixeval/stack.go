package infix

import (
	"sync"
)

// Stack implements generic stack
type Stack struct {
	sync.Mutex
	item []interface{}
}

// Peek shows the top element of the stack without removing it from the stack
func (s *Stack) Peek() interface{} {
	s.Lock()
	defer s.Unlock()

	if len(s.item) == 0 {
		return ""
	}

	return s.item[len(s.item)-1]
}

// Push inserts new item onto top of the stack
func (s *Stack) Push(x interface{}) {
	s.Lock()
	defer s.Unlock()
	s.item = append(s.item, x)
}

// Pop pops the top item of the stack
func (s *Stack) Pop() (x interface{}) {
	s.Lock()
	defer s.Unlock()
	x, s.item = s.item[len(s.item)-1], s.item[:len(s.item)-1]
	return
}

// IsEmpty checks if given stack is empty
func (s *Stack) IsEmpty() bool {

	if len(s.item) == 0 {
		return true
	}
	return false
}

// NewStack creates and returns new stack
func NewStack() (s *Stack) {
	s = &Stack{}
	return
}
