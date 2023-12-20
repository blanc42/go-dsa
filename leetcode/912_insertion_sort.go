package leetcode

func SortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				temp := nums[j-1]
				nums[j-1] = nums[j]
				nums[j] = temp
			} else {
				break
			}
		}
	}
	return nums
}
