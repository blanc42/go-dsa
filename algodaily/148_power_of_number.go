package algodaily

func PowerOfNumber(x, n int) float64 {
	if n == 0 {
		return 1
	} else if n < 0 {
		var y float64 = 1
		for i := 0; i < n; i++ {
			y *= float64(x)

		}
		return 1 / y
	} else {
		var y float64 = 1
		for i := 0; i < n; i++ {
			y *= float64(x)
		}
		return y
	}
}
