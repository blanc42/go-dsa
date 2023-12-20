package leetcode

func RemoveElement(nums []int, val int) int {
	s := 0 // swaps
	l := len(nums)
	i := 0
	for i <= l-s-1 {
		if nums[i] == val {
			temp := nums[i]
			nums[i] = nums[l-s-1]
			nums[l-s-1] = temp
			s++
			continue
		}
		i++
	}
	return l - s
}
