package leetcode

func RemoveDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	read := 1
	write := 1

	for read < len(nums) {
		if nums[read] != nums[write-1] {
			nums[write] = nums[read]
			write++
		}
		read++
	}

	// if len(nums) < 2 {
	//   return len(nums)
	// }
	// now := 0
	// next := 1
	// for next < len(nums) {
	//   if nums[now] == nums[next] {
	//     nums[now] = nums[next]
	//   }
	//   now++
	//   next++
	// }
	// return now + 1
	return write
}
