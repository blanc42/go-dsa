package leetcode

func GeneratePascalTriangle(numRows int) [][]int {
	out := [][]int{{1}}
	if numRows == 1 {
		return out
	}
	out = append(out, []int{1, 1})
	start := 3
	curr := []int{}
	i := 0
	for start <= numRows {
		if i == 0 || i == start-1 {
			curr = append(curr, 1)
		} else {
			l := len(out) - 1
			last := out[l]
			curr = append(curr, last[i]+last[i-1])
		}
		if i == start-1 {
			out = append(out, curr)
			curr = []int{}
			start++
			i = 0
			continue
		}
		i++
	}
	return out
}
