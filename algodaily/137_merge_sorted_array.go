package algodaily

func MergeSortedArrary(A, B []int) []int {
	var C []int
	len_a := len(A)
	len_b := len(B)

	i, j := 0, 0
	for ; i < len_a; i++ {
		for j < len_b {
			if B[j] < A[i] {
				C = append(C, B[j])
				j++
			} else {
				break
			}
		}
		C = append(C, A[i])
	}
	C = append(C, B[j:]...)
	return C

}
