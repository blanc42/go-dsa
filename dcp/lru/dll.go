package lur

// doubly linked list for LRU implementations

type Node struct {
	value    int
	previous *Node
	next     *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (list *DoublyLinkedList) New() *DoublyLinkedList {
	d := new(DoublyLinkedList)
	return d
}

func (list *DoublyLinkedList) Add(node *Node) {

}

func (list *DoublyLinkedList) Update(node *Node) {
	node.value += 1
	if list.head == node {
		return
	}

}

// sawps two nodes
// todo : check for null pointer exceptions
func MoveUp(node *Node) {
	A := node.previous
	A0 := node.previous.previous
	B := node
	B0 := node.next

	// moving B up
	A0.next = B
	B.previous = A0

	// moving A down
	A.next = B0
	B0.previous = A

	// connecting A, B
	B.next = A
	A.previous = B

}
