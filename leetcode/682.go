package leetcode

import "strconv"

func CalPoints(operations []string) int {
	var out []int
	for _, op := range operations {
		if num, err := strconv.Atoi(op); err == nil {
			out = append(out, num)
		} else {
			l := len(out)
			switch op {
			case "+":
				out = append(out, out[l-1]+out[l-2])
			case "D":
				out = append(out, 2*out[l-1])
			case "C":
				out = out[:l-1]
			}
		}
	}
	sum := 0
	for _, num := range out {
		sum += num
	}
	return sum
}
