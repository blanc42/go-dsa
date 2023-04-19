package dll

import "fmt"

// type doubly linked list
type Dll struct {
	Len   int
	Front *Node
	Tail  *Node
}

// node
type Node struct {
	Data int
	Prev *Node
	Next *Node
}

func (dl *Dll) Append(data int) {
	newNode := new(Node)
	newNode.Data = data
	if dl.Len == 0 {
		dl.Front = newNode
		dl.Tail = newNode
	} else {
		newNode.Prev = dl.Tail
		dl.Tail.Next = newNode
		dl.Tail = newNode
	}
	dl.Len++
}

func (dl *Dll) Print() {
	if dl.Len == 0 {
		fmt.Println("empty list")
	} else {
		current := dl.Front
		for {
			fmt.Print(current.Data)
			if current.Next == nil {
				fmt.Println("")
				break
			} else {
				fmt.Print(" -> ")
				current = current.Next
			}
		}
	}
}



