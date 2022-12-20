package utils

func ArrayContains(a []int, b int) bool {
	for _, x := range a {
		if x == b {
			return true
		}
	}
	return false
}
