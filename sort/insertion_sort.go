package sort

import "dsa/utils"

func InsertionSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {

		curr := i
		for curr > 0 && arr[curr] < arr[curr-1] {
			utils.SwapInt(&arr[curr-1], &arr[curr])
			curr--
		}
	}

	return arr
}
