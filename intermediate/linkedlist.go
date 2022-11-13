package main

import "fmt"

type node struct {
	value int
	next  *node
}

type val int

func main() {
	fmt.Println("Inside main")
	ll := NewLinkedList()
	fmt.Println("Empty linked list:", ll, ll.value, ll.next)
	for i := 1; i < 10; i++ {
		ll = ll.Insert(val(i))
	}
	ll.Display()
}

// NewLinkedList() creates and returns a nil linked list (i.e. *node).
func NewLinkedList() *node {
	return &node{}
}

// Insert() inserts the value v at the beginning of the linked list
func (n *node) Insert(v val) *node {
	return &node{value: int(v), next: n}
}

// Append() inserts the value v at the end of the linked list
func (n *node) Append(v val) *node {
	return &node{value: int(v), next: n}
}

// Display() prints all the elements of the linked list starting from beginning
func (n *node) Display() {
	fmt.Print("Inside Display.\n First node::")
	for n != nil {
		fmt.Println(n.value)
		fmt.Print("Next node:: ")
		n = n.next
	}
}

// Pop() pops(removes) the head element from current linked list and returns the updated linked list
// func (n *node) Pop() *node

// RemoveLast() removes the tail(last) element from current linked list and returns the updated linked list
// func (n *node) RemoveLast() *node

// Remove() removes all the nodes with data==v from current linked list and returns the updated linked list
// func (n *node) Remove(v val) *node
