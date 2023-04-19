/*
this file contains source code for a singly linked list data structure
*/
package ll

// node struct

/*
STRUCTS
-------
node struct { data, next pointer }
Linkedlist struct { head pointer, length, tail pointer }

FUNCTIONS
----------
New() creates and returns a Linkedlist


METHODS
-------
Append(key int)
Prepend(key int)
Insert(pos int, key int)

Extract(pos int)
Pop() last element
Deqeue() first element

Update(pos int, key int)

Delete(pos int)

Print() prints the linked list in a pretty format
 - this one should be String methods so the when used with Prinln it can be called

*/

type node struct {
	data int
	next *node
}

type Linkedlist struct {
	head *node
	len  int
	tail *node
}

func (l *Linkedlist) Append(key int) {
	n := &node{key, nil}
	l.tail.next = n
	l.tail = n
}

func (l *Linkedlist) Print() {
	// current := h.head
}
