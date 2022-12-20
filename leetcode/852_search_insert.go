package leetcode

func SearchInsert(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for {
		if start <= end {
			i := (start + end) / 2
			mid := nums[i]
			if mid == target {
				return i
			} else if mid < target {
				start = i + 1
			} else {
				end = i - 1
			}
		} else {
			return start
		}
	}
}
