package sort

import "dsa/algodaily"

// not working yet
// func merge(a []int, b []int) {
// 	lena, lenb := len(a), len(b)
// 	for i, j := 0, 0; i < lena && j < lenb; i++ {
// 		if a[i] > b[j] {
// 			utils.SwapInt(&a[i], &b[j])
// 			j++
// 		}
// 	}
// }

// func MergeSortedArrarys(A,B []int) []int {
// 	var C []int

// 	return C
// }

// func MergeSort(nums []int) {
// 	len := len(nums)
// 	mid := len / 2
// 	if len >= 2 {
// 		MergeSort(nums[0 : mid-1])
// 		MergeSort(nums[mid:len])
// 		merge(nums[0:mid-1], nums[mid:len])
// 	}
// }

func MergeSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	left := MergeSort(nums[:n/2])
	right := MergeSort(nums[n/2:])

	ans := algodaily.MergeSortedArrary(left, right)
	return ans
}
