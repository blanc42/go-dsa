package leetcode

import (
	"math"
	"sort"
)

func NumSubseq(nums []int, target int) int {
	sort.Ints(nums)
	start := 0
	end := len(nums) - 1
	var total float64
	mod := math.Pow10(9) + 7
	for start <= end {
		if nums[start]+nums[end] > target {
			end--
		} else {
			total = math.Mod(total+math.Pow(2, float64(end-start)), mod)
			start++
		}
	}
	return int(total)
}
