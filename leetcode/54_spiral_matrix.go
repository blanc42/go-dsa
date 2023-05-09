package leetcode

func SpiralOrder(matrix [][]int) []int {
	var out []int
	length := len(matrix)

	for length > 0 {
		// append the firs row
		out = append(out, matrix[0]...)

		matrix = matrix[1:]
		length = len(matrix)
		if length == 0 {
			return out
		}

		// append the last element of all the other rows and update the row
		for j := 0; j < length; j++ {
			l := len(matrix[j])
			out = append(out, matrix[j][l-1])
			matrix[j] = matrix[j][:l-1]
		}
		if len(matrix[0]) == 0 {
			return out
		}

		// append the last row
		// reverse the last row
		for i, j := 0, len(matrix[length-1])-1; i < j; i, j = i+1, j-1 {
			matrix[length-1][i], matrix[length-1][j] = matrix[length-1][j], matrix[length-1][i]
		}
		out = append(out, matrix[length-1]...)

		matrix = matrix[:length-1]
		length = len(matrix)
		if length == 0 {
			return out
		}

		// append the first element from the other row excluding the first and last
		for j := length - 1; j >= 0; j-- {
			out = append(out, matrix[j][0])
			matrix[j] = matrix[j][1:]
		}
		if len(matrix[0]) == 0 {
			return out
		}
	}
	return out
}
