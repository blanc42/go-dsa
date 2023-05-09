package leetcode

import "math"

func Atoi(s string) int {

	// if len(s) == 0 {
	//     return 0
	// }
	neg := 0
	var start int
	for _, c := range s {
		r := rune(c)
		if r == ' ' {
			start++
		} else if r == '-' && neg != 0 {
			return 0
		} else if r == '+' && neg != 0 {
			return 0
		} else if r == '-' {
			neg = -1
			start++
		} else if r == '+' {
			neg = 1
			start++
			break
		} else {
			if neg == 0 {
				neg = 1
			}
			break
		}
	}
	s = s[start:]

	var out uint
	for _, c := range s {
		b := byte(c)
		val := uint(b - '0')
		if val < 0 || val > 9 {
			return re(out, neg)
		}
		out = out*10 + val
	}

	return re(out, neg)
}

func re(out uint, neg int) int {
	modPos := int(math.Pow(2, 31)) - 1
	modNeg := -1 * int(math.Pow(2, 31))
	normalout := int(out) * neg
	if out > uint(modPos) && neg == 1 {
		return modPos
	}
	if out > uint(modNeg*-1) && neg == -1 {
		return modNeg
	}
	return normalout
}
