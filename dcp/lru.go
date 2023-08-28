package dcp

// least recently used cache problem
// daily coding problem 697

// type Pair struct {
// 	value int
// 	node  *Node
// }

// type LRUcache struct {
// 	Max   int
// 	Count int
// 	Cache map[int]Pair
// 	Dll   *DoublyLinkedList
// }

// // initializer a new LRU cache
// func (l *LRUcache) New(max int) {
// 	l.Max = max
// 	l.Cache = make(map[int]Pair)
// 	l.Dll = l.Dll.New()
// }

// // get the vlues if exists and return null if does not
// func (l *LRUcache) Get(key int) int {
// 	if pair, ok := l.Cache[key]; ok {
// 		l.Dll.Update(pair.node)
// 		return pair.value
// 	} else {
// 		return 0
// 	}

// }

// // sets the key to a value
// func (l *LRUcache) Set(key int, value int) {
// 	newNode := new(Node)
// 	l.Dll.Add(newNode)

// }

//  ============

// type Node struct {
// 	value    int
// 	previous *Node
// 	next     *Node
// }

// type DoublyLinkedList struct {
// 	head *Node
// 	tail *Node
// }

// func (list *DoublyLinkedList) New() *DoublyLinkedList {
// 	d := new(DoublyLinkedList)
// 	return d
// }

// func (list *DoublyLinkedList) Add(node *Node) {

// }

// func (list *DoublyLinkedList) Update(node *Node) {
// 	node.value += 1
// 	if list.head == node {
// 		return
// 	}

// }

// // sawps two nodes
// // todo : check for null pointer exceptions
// func MoveUp(node *Node) {
// 	A := node.previous
// 	A0 := node.previous.previous
// 	B := node
// 	B0 := node.next

// 	// moving B up
// 	A0.next = B
// 	B.previous = A0

// 	// moving A down
// 	A.next = B0
// 	B0.previous = A

// 	// connecting A, B
// 	B.next = A
// 	A.previous = B

// }
