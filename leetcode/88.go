package leetcode

func Merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	last := m + n - 1
	for last >= 0 {
		if j < 0 {
			break
		}
		if i < 0 {
			nums1[last] = nums2[j]
			j--
		} else {
			if nums1[i] > nums2[j] {
				nums1[last] = nums1[i]
				i--
			} else {
				nums1[last] = nums2[j]
				j--
			}
		}
		last--
	}
}
