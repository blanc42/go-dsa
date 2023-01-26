package leetcode

func ContainsDuplicate(nums []int) bool {
	var a = make(map[int]int)
	for _, n := range nums {
		if v := a[n]; v == 1 {
			return true
		} else {
			a[n] = 1
		}
	}
	return false
}
