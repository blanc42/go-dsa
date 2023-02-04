/*
this file contains source code for a singly linked list data structure
*/
package ll

import "fmt"

type Node struct {
	Data int
	Next *Node
}

// default stuct node
type LinkedList struct {
	head *Node
}

// add a node to the start of the list
func (l *LinkedList) Prepend(n *Node) {
	n.Next = l.head
	l.head = n
}

// add a note at the end of the list
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

// delete a given node from the list 
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

// update a part of the node, takes in old node and new node
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


// print the linked list in 'x -> y -> z' format
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
