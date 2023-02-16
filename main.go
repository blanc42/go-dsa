package main

import (
	"dsa/ds/heap"
	"fmt"
)

func main() {
	log := fmt.Println
	// A := []int{2, 1, 3, 4, 5, 9, 7, 6, 8}
	h := &heap.MaxHeap{}
	h.Insert(10)
	fmt.Println(h)
	h.Insert(12)
	fmt.Println(h)
	h.Insert(1)
	fmt.Println(h)
	h.Insert(14)
	fmt.Println(h)
	// extrcting
	log(h.Extract())
	fmt.Println(h)
	// extrcting
	log(h.Extract())
	fmt.Println(h)
	// extrcting
	log(h.Extract())
	fmt.Println(h)

}
