package ds

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Prepend(n *Node) {
	n.Next = l.head
	l.head = n
}

func (l *LinkedList) AddNode(n *Node) {
	n.Next = nil
	if l.head == nil {
		l.head = n
	} else {
		iter := l.head
		for iter.Next != nil {
			iter = iter.Next
		}
		iter.Next = n
	}

}

func (l *LinkedList) DeleteNode(n *Node) {
	if l.head.Data == n.Data {
		l.head = n.Next
	}
	var prevNode *Node = l.head
	current := l.head.Next

	for {
		if current != nil {
			if current.Data == n.Data {
				prevNode.Next = current.Next
				break
			} else {
				prevNode = current
				current = prevNode.Next
			}
		} else {
			fmt.Println("no such Data exist")
			break
		}
	}

}

func (l *LinkedList) UpdateNode(old *Node, new *Node) {
	iter := l.head
	for {
		if iter != nil {
			if iter.Data == old.Data {
				iter.Data = new.Data
			} else {
				iter = iter.Next
			}
		} else {
			break
		}
	}
}

func (l *LinkedList) Print() {
	if l.head == nil {
		fmt.Println("list is empty")
	} else {
		temp := l.head
		for {
			if temp.Next != nil {
				fmt.Print(temp.Data)
				fmt.Print(" -> ")
				temp = temp.Next
			} else {
				fmt.Print(temp.Data)
				fmt.Println()
				break
			}
		}
	}
}
