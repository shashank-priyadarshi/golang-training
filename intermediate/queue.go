package main

import "fmt"

type queue []int

func main() {
	fmt.Print("Inside main\n")
	q := NewQueue()
	fmt.Println("Queue at creation:: ", q)
	fmt.Println("Is queue empty at creation:: ", q.Empty()) // true
	fmt.Println("Length of queue at creation:: ", q.Size()) // 0
	q = q.Push(1)
	q = q.Push(2)
	q = q.Push(3)
	q = q.Push(4)
	q = q.Push(5)
	fmt.Println("Queue after pushing 5 elements:: ", q) // [1 2 3 4 5]
	fmt.Println("Is queue empty:: ", q.Empty())         // false
	fmt.Println("Size of queue:: ", q.Size())           // 5
	fmt.Println("First element in queue:: ", q.Front()) // 1
	fmt.Println("Last element in queue:: ", q.Back())   // 5
	q = q.Pop()
	fmt.Println("First element after pop operation:: ", q.Front()) // 2
	fmt.Println("Last element after pop operation:: ", q.Back())   // 5
	fmt.Println("Queue:: ", q)                                     // [2 3 4 5]
	fmt.Println("Exiting main")
}

// NewQueue() creates an empty queue and returns it.
func NewQueue() queue {
	return queue{}
}

// q.Empty() returns true if it is empty
func (q queue) Empty() bool {
	return len(q) == 0
}

// q.Size() returns the total number of elements in q
func (q queue) Size() int {
	return len(q)
}

// q.Front() returns the element present at the front of the queue q (i.e. the element which came first in the queue)
func (q queue) Front() int {
	return q[0]
}

// q.Back() returns the element present at the back of the queue q (i.e. the element which came last in the queue)
func (q queue) Back() int {
	return q[len(q)-1]
}

// q.Push(v) pushes a new value at the end(back) of current q and returns the updated queue
func (q queue) Push(v int) queue {
	return append(q, v)
}

// q.Pop() pops(removes) the front value from current q and returns the updated queue
func (q queue) Pop() queue {
	return q[1:]
}
