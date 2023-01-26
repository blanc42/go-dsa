package algodaily

import "dsa/utils"

func ArraryIntersection(a []int, b []int) []int {
	var output []int

	for _, x := range a {
		if !utils.ArrayContains(output, x) {
			for _, y := range b {
				if x == y {
					output = append(output, x)
					break
				}
			}
		}
	}

	return output
}
