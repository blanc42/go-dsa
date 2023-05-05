package sort

import (
	"dsa/utils"
)

func BubbleSort(arr []int) []int {
	n := len(arr)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				utils.SwapInt(&arr[i], &arr[j])
			}
		}
	}
	return arr
}
