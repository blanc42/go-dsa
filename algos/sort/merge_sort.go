package sort

import "dsa/utils"

// not working yet

func merge(a []int, b []int) {
	lena, lenb := len(a), len(b)
	for i, j := 0, 0; i < lena && j < lenb; i++ {
		if a[i] > b[j] {
			utils.Swap(&a[i], &b[j])
			j++
		}
	}

}

func MergeSort(nums []int) {
	len := len(nums)
	mid := len / 2
	if len >= 2 {
		MergeSort(nums[0 : mid-1])
		MergeSort(nums[mid:len])
		merge(nums[0:mid-1], nums[mid:len])
	}
}
