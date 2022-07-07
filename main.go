package main

import (
	"fmt"
)

// Node presents one element of stack
type Node struct {
	prev *Node
	Data interface{}
}

// Stack presents classic LIFO stack
type Stack struct {
	top    *Node
	length int
}

// Push add new element to end
func (s *Stack) Push(el *Node) {
	if el == nil {
		return
	}

	el.prev = s.top
	s.top = el

	s.length++
}

// emptyStack presents empty stack
const emptyStack = "[]"

// String returns string of all elements
func (s Stack) String() string {
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
func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}
	s.top = s.top.prev
	s.length--
}

// IsEmpty returns result of empty stack check
func (s Stack) IsEmpty() bool {
	return s.length == 0
}

func main() {
	var stack Stack

	el1 := Node{Data: "!"}
	el2 := Node{Data: "world"}
	el3 := Node{Data: "Hello"}

	stack.Push(&el1)
	stack.Push(&el2)
	stack.Push(&el3)

	fmt.Printf("STACK WITH 3 ELEMENTS: %+v\n", stack)
	fmt.Printf("STACK WITH IS EMPTY: %+v\n", stack.IsEmpty())

	stack.Pop()
	stack.Pop()
	fmt.Printf("STACK WITH 1 ELEMENT: %+v\n", stack)
	stack.Pop()
	fmt.Printf("STACK WITH IS EMPTY: %+v\n", stack.IsEmpty())

}
