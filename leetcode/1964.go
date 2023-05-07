package leetcode

// TLE error
func LongestObstacleCourseAtEachPosition(obstacles []int) []int {
	var out []int
	out = append(out, 1)
	l := len(obstacles)
	for i := 1; i < l; i++ {
		var max int
		for b := i - 1; b >= 0; b-- {
			if obstacles[b] <= obstacles[i] && out[b] > max {
				max = out[b]
			}
		}
		out = append(out, max+1)
	}
	return out
}
