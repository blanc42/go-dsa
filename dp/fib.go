package dp

func Fibonacci(x int) uint64 {
	memo := make(map[int]uint64)
	memo[0] = 0
	memo[1] = 1

	var fibonacciHelper func(int) uint64
	fibonacciHelper = func(n int) uint64 {
		if val, ok := memo[n]; !ok {
			result := fibonacciHelper(n-1) + fibonacciHelper(n-2)
			memo[n] = result
			return result
		} else {
			return val
		}
	}

	return fibonacciHelper(x)
}
