package leetcode

func DiagonalSum(mat [][]int) int {
	var sum int
	len := len(mat[0])
	for i, a := range mat {
		if i == len-1-i {
			sum += a[i]
		} else {
			sum += a[i] + a[len-1-i]
		}
	}
	return sum
}
