package ds

import "fmt"

type Node[T interface{}] struct {
	value    T
	previous *Node[T]
	next     *Node[T]
}

func NewNode[T interface{}](val T) *Node[T] {
	node := &Node[T]{value: val}
	return node
}

type Deque[T interface{}] struct {
	head *Node[T]
	tail *Node[T]
}

func NewDeque[T comparable]() *Deque[T] {
	return &Deque[T]{}
}

// checks if the head is nil, which means the deque is empty
func (d *Deque[T]) IsEmpty() bool {
	return d.head == nil
}

// push at the end of the stack
func (d *Deque[T]) Push(value T) {
	node := NewNode(value)
	// checking if there is no elements in the stack
	if d.tail == nil {
		d.head = node
		d.tail = node
		return
	}
	node.previous = d.tail
	d.tail.next = node
	d.tail = node
}

// pop from the end of the stack
func (d *Deque[T]) Pop() (value T) {
	if d.tail == nil {
		return Node[T]{}.value
	}
	if d.head == d.tail {
		val := d.tail.value
		d.head = nil
		d.tail = nil
		return val
	}
	node := d.tail
	d.tail = node.previous
	d.tail.next = nil
	return node.value
}

// enqueue at the end of the queue
func (d *Deque[T]) Enqueue(value T) {
	d.Push(value)
}

func (d *Deque[T]) EnqueueSlice(values []T) {
	for _, val := range values {
		d.Push(val)
	}
}

// dequeue at the start of the queue
func (d *Deque[T]) DeQueue() (value T) {
	if d.head == nil {
		return Node[T]{}.value
	}
	if d.head == d.tail {
		val := d.head.value
		d.head = nil
		d.tail = nil
		return val
	}
	node := d.head
	d.head = node.next
	d.head.previous = nil
	return node.value
}

// print the dequeu
func (d *Deque[T]) Print() {
	curr := d.head
	for curr != nil {
		fmt.Print(curr.value)
		if curr.next != nil {
			fmt.Print(" <-> ")
		}
		curr = curr.next
	}
	fmt.Println("")
}
