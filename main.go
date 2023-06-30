package main

import "fmt"

func main() {
	// ll := dcp.NewLinkedList()

	// ll.Push(&dcp.Node{Value: 7})
	// ll.Push(&dcp.Node{Value: 7})
	// ll.Push(&dcp.Node{Value: 3})
	// ll.Push(&dcp.Node{Value: 5})
	// // ll.Push(&dcp.Node{Value: 5})

	// ll.Print()

	// ll.MoveByK(2)
	// ll.Print()

	g := Graph{}
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 4)
	g.Print()

	d := Deque[int]{}
	d.Push(10)
	d.Print()
	d.Push(11)
	d.Print()
	d.Push(12)
	d.Print()
	d.Push(13)
	fmt.Println(d.DeQueue())
	d.Print()
	d.Enqueue(14)
	d.Print()
	fmt.Println(d.DeQueue())
	d.Print()
	d.Enqueue(15)
	fmt.Println(d.DeQueue())
	fmt.Println(d.DeQueue())
	d.Print()
}
