package ds

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	count int
	head  *Node
}

func (l *LinkedList) Prepend(n *Node) {
	n.next = l.head
	l.head = n
}

// add node
func (l *LinkedList) AddNode(n *node) {

}
func (l *LinkedList) DeleteNode(n *node) {

}
func (l *LinkedList) UpdateNode(n *node) {

}
func (l *LinkedList) Print(n *node) {
	temp := l.head
	for temp != nil {
		fmt.Println(temp.data, " -> ")
	}
	fmt.Println()
}

// function to create linked list
// func CreateList() (h *node) {
// 	var head *node
// 	head = nil
// 	return head
// }

// method to add a node
// adds a node at the end of the list
// func (head *node) AddNode(d int) {
// 	// temp := head
// }

// var a []int = []int{1, 2, 3}

// method to print the linked list
// func (h *node) String() {
// 	temp := h
// 	for temp.next != nil {
// 		fmt.Printf("%v ->", temp.data)
// 		temp = temp.next
// 	}
// }

// method to delete a node if exists
// func (h *node) DeleteNode(d int) {

// 	var prevNode *node
// 	var currentNode *node

// 	we loop over the list and if check the data if each node
// 	if any node matches then we take the past node's next pointer to the current node's next pointer

// }

// method to update a node

// insert node

// func main() {
// 	// var head *node
// 	// head = nil
// 	head := createList()
// 	temp := new(node)
// 	(temp).data = 12
// 	head = temp
// 	// head.addNode(12)
// 	// head.addNode(14)
// 	fmt.Println(head)
// }
