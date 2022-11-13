package main

import "fmt"

type stack struct {
	top       int
	container []int
}

func main() {
	fmt.Println("Inside main")
	s := NewStack()
	fmt.Println("Empty stack:", s)
	fmt.Println("Is stack empty:", s.Empty())
	for i := 1; i < 10; i++ {
		s = s.Push(i)
	}
	fmt.Println("Stack after pushing 10 elements:", s)
	fmt.Println("Is stack empty:", s.Empty())
	fmt.Println("Size of stack:", s.Size())
	fmt.Println("Top of stack:", s.Top())
	s = s.Pop()
	fmt.Println("Stack after popping:", s)
}

// NewStack() creates and returns a stack with top = -1 and empty container.
func NewStack() stack {
	return stack{top: -1, container: []int{}}
}

// s.Empty() returns true if top == -1
func (s stack) Empty() bool {
	return s.top == -1
}

// s.Top() returns the value at the top of container (the value which was added most recently to the container)
func (s stack) Top() int {
	return s.container[s.top]
}

// s.Size() returns the total elements present in the container
func (s stack) Size() int {
	return s.top + 1
}

// s.Push(v) pushes a new value at the top of current s and returns the updated stack
func (s stack) Push(v int) stack {
	s.top++
	s.container = append(s.container, v)
	return s
}

// s.Pop() pops(removes) the top value from current s and returns the updated stack
func (s stack) Pop() stack {
	s.container = s.container[:s.top]
	s.top--
	return s
}
