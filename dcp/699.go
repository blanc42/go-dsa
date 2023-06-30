package dcp

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Length int
	Head   *Node
}

func NewLinkedList() *LinkedList {
	l := LinkedList{}
	return &l
}

func (l *LinkedList) Push(node *Node) {
	if l.Head == nil {
		l.Head = node
		l.Length++
		return
	}
	curr := l.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = node
	l.Length++
}

func (l *LinkedList) Print() {
	curr := l.Head

	for curr.Next != nil {
		fmt.Print(curr.Value, "->")
		curr = curr.Next
	}
	fmt.Println(curr.Value, "->nil")
}

func (l *LinkedList) MoveByK(k int) {
	head := l.Head
	tail := head
	// finding the tail of the first part
	for i := 0; i < l.Length-k-1; i++ {
		tail = tail.Next
	}

	if tail.Next != nil {
		curr := tail.Next
		tail.Next = nil
		l.Head = curr

		// attaching the both lists (trying to find the head of secnond list)
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = head
	}
}
