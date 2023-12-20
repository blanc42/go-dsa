package main

import (
	"dsa/dp"
	"dsa/utils"
)

func main() {
	utils.Log("==== Golang ====")

	for i := 0; i < 46; i++ {

		utils.Log(dp.Fibonacci(i))
	}

}

// func fibonacci64(n int) *big.Int {
// 	a, b := big.NewInt(0), big.NewInt(1)

// 	for i := 0; i < n; i++ {
// 		a.Add(a, b)
// 		a, b = b, a
// 	}

// 	return a
// }

// fibonicci :DONE:
// bfs
// dfs
// dijkstra
// naieve string ssearch
// kmp
// longest common subsequence
// robin karp
// number of islands
// preorder, postorder, in order traversals

// Longest increasing subsequence
// tower of hanoi
// coin change I, coin change II
