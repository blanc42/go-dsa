package sort

import (
	"dsa/utils"
	"fmt"
)

func SelectionSort(arr []int) []int {
	fmt.Println(arr)
	l := len(arr)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if arr[i] > arr[j] {
				utils.Swap(&arr[i], &arr[j])
			}
		}
	}
	return arr
}
