package sort

import (
	"dsa/utils"
)

func SelectionSort(arr []int) []int {
	l := len(arr)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if arr[i] > arr[j] {
				utils.SwapInt(&arr[i], &arr[j])
			}
		}
	}
	return arr
}
