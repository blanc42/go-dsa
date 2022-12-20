package algos

import (
	"dsa/utils"
)

func BubbleSort(arr []int) []int {
	var swapped bool = true
	l := len(arr)
	for swapped {
		swapped = false
		for i := 0; i < l-1; i++ {
			if arr[i] > arr[i+1] {
				utils.Swap(&arr[i], &arr[i+1])
				swapped = true
			}
		}
	}
	return arr
}
