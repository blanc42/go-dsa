package algodaily

func LonelyNumber(A []int) int {
	M := make(map[int]int)
	for _, x := range A {
		if _, ok := M[x]; ok {
			delete(M, x)
		} else {
			M[x]++
		}
	}
	var out int
	for k, _ := range M {
		out = k
	}
	return out
}
