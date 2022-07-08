package main

import (
	"fmt"
)

// Node presents one element of stack
type Node[T any] struct {
	prev *Node[T]
	Data T
}

// Stack presents classic LIFO stack
type Stack[T any] struct {
	top    *Node[T]
	length int
}

// Push add new element to end
func (s *Stack[T]) Push(el T) {
	n := &Node[T]{Data: el}

	n.prev = s.top
	s.top = n

	s.length++
}

// emptyStack presents empty stack
const emptyStack = "[]"

// String returns string of all elements
func (s Stack[T]) String() string {
	if s.IsEmpty() {
		return emptyStack
	}

	result := fmt.Sprintf("[%v", s.top.Data)
	current := s.top
	for current.prev != nil {
		result += fmt.Sprintf(", %v", current.prev.Data)
		current = current.prev
	}
	result += "]"

	return result
}

// Pop remove one element from prev
func (s *Stack[T]) Pop() T {
	val := *new(T)
	if s.IsEmpty() {
		return val
	}
	val = s.top.Data
	s.top = s.top.prev
	s.length--

	return val
}

// IsEmpty returns result of empty stack check
func (s Stack[T]) IsEmpty() bool {
	return s.length == 0
}

func main() {
	var stack Stack[string]

	stack.Push("one")
	stack.Push("two")
	stack.Push("three")

	fmt.Printf("STACK WITH 3 ELEMENTS: %+v\n", stack)
	fmt.Printf("STACK WITH IS EMPTY: %+v\n", stack.IsEmpty())

	el1 := stack.Pop()
	fmt.Printf("STACK RETURNED ONE ELEMENT: %+v\n", el1)
	stack.Pop()
	fmt.Printf("STACK WITH 1 ELEMENT: %+v\n", stack)
	stack.Pop()
	fmt.Printf("STACK WITH IS EMPTY: %+v\n", stack.IsEmpty())
}
