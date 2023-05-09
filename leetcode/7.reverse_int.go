package leetcode

import "math"

func Reverse(x int) int {
	var out int
	curr := x
	modNeg := -1 * int(math.Pow(2, 31))
	modPos := int(math.Pow(2, 31)) - 1
	for curr%10 == 0 && curr != 0 {
		curr = curr / 10
	}
	for curr != 0 {
		rem := curr % 10
		out = out*10 + rem
		curr = curr / 10
		if out-modNeg < 0 || out-modPos > 0 {
			return 0
		}
	}
	return out
}
