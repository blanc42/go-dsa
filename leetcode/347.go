package leetcode

import "sort"

func TopKFrequent(nums []int, k int) []int {

	// sort the array
	sort.Ints(nums)

	l := len(nums)
	// create a new slice with the length of nums
	countSlice := make([][]int, 0) // Create an empty slice of slices

	// Append inner slices as needed
	for i := 0; i < l; i++ {
		innerSlice := []int{}                       // Create an empty inner slice
		countSlice = append(countSlice, innerSlice) // Append the inner slice to the outer slice
	}

	curr := nums[0]
	count := 0
	for _, v := range nums {
		if v == curr {
			count++
		} else {
			countSlice[count-1] = append(countSlice[count-1], curr)
			curr = v
			count = 1
		}
	}
	// add the last element
	countSlice[count-1] = append(countSlice[count-1], curr)

	out := []int{}

	for i := l - 1; i >= 0; i-- {
		if len(out) == k {
			break
		}
		out = append(out, countSlice[i]...)
	}

	return out
}
