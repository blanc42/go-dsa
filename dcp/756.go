package dcp

// This problem was asked by Slack.
// You are given a string formed by concatenating several words corresponding to the integers zero through nine and then anagramming.
// For example, the input could be 'niesevehrtfeev', which is an anagram of 'threefiveseven'. Note that there can be multiple instances of each integer.
// Given this string, return the original integers in sorted order. In the example above, this would be 357.

func AnaStringToInt(s string) int {
	res := 0

	// letters vector
	lvec := make(map[rune]int)

	// convert the s into the letter vector
	for _, r := range s {
		lvec[r]++
	}

	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	// range over the numbers and
	for i, number := range numbers {
		for check(number, lvec) {
			res = res*10 + i + 1
		}
	}
	return res
}

// check if that give number exist in the arry
func check(number string, lv map[rune]int) bool {
	for _, r := range number {
		if lv[r] == 0 {
			return false
		}
	}
	for _, r := range number {
		lv[r]--
	}
	return true
}
