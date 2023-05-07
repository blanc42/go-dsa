package leetcode

func maxVowels(s string, k int) int {

	// calculating the inital sum and it will act as maxSum
	var sum int
	for _, c := range s[:k] {
		sum += checkVowel(string(c))
	}

	// present sum
	currSum := sum

	// looping thorugh the remaining part and finding max sum
	for i := k; i < len(s); i++ {
		currSum = currSum - checkVowel(string(s[i-k])) + checkVowel(string(s[i]))
		if currSum > sum {
			sum = currSum
		}
	}
	return sum

}

func checkVowel(c string) int {
	if c == "a" || c == "e" || c == "i" || c == "o" || c == "u" {
		return 1
	}
	return 0
}
