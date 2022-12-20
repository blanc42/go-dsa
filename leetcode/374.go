package leetcode

func guess(guess int) int {
	pick := 13
	if guess == pick {
		return 0
	} else if guess < pick {
		return 1
	} else {
		return -1
	}
}

func GuessNumber(n int) int {
	start := 1
	end := n
	for {
		mid := (start + end) / 2
		guess := guess(mid)
		switch guess {
		case 0:
			return mid
		case 1:
			start = mid + 1
		case -1:
			end = mid - 1
		}
	}
}


