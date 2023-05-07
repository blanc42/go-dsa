package leetcode

func LengthOfLongestSubstring(s string) int {
	m := make(map[byte]struct{})
	max := 0
	start := 0

	for end := 0; end < len(s); end++ {
		// check if the rune exits in map
		if _, ok := m[s[end]]; ok {
			// if exists then move the start till it becomes unique again and delete the repeated one from map
			for start <= end {
				delete(m, s[start])
				if s[start] == s[end] {
					start++
					break
				}
				start++
			}
		}
		// add the end to the map
		m[s[end]] = struct{}{}

		// update the max
		curr_max := end - start + 1
		if curr_max > max {
			max = curr_max
		}
	}
	return max
}
