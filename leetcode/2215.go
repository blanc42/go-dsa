package leetcode

func FindDifference(nums1 []int, nums2 []int) [][]int {
	m1 := make(map[int]struct{})
	for _, v := range nums1 {
		m1[v] = struct{}{}
	}
	m2 := make(map[int]struct{})
	for _, v := range nums2 {
		if _, ok := m1[v]; !ok {
			m2[v] = struct{}{}
		}
	}

	var s1, s2 []int
	for k, _ := range m1 {
		s1 = append(s1, k)
	}
	for k, _ := range m2 {
		s2 = append(s2, k)
	}
	return [][]int{s1, s2}

}
