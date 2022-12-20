package main

import (
	"dsa/algodaily"
	"fmt"
)

func main() {
	// fmt.Println(algodaily.StringReverse("something is bad"))
	b := []int{-1, 0, 3, 5, 9, 12}
	a := []int{11, 12, 13, 5, 6}
	// fmt.Println(leetcode.BinarySearch(a, 3))
	// fmt.Println(leetcode.GuessNumber(16))
	// fmt.Println(searchInsert(a, 5))
	// fmt.Println(a)
	// algos.MergeSort(a)
	// fmt.Println(a)
	fmt.Println(algodaily.ArraryIntersection(a, b))
}
