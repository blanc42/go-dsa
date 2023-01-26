package algodaily

import "strings"

// Reverse a string Algodaily
func StringReverse(i string) (output string) {
	// split the string into characters
	input := strings.Split(i, "")
	// range over the slice and swap the first character with last character until i < j
	j := len(input)
	for i := 0; i < j; i++ {
		temp := input[i]
		input[i] = input[j-1]
		input[j-1] = temp
		j--
	}
	// join the characters in slice
	output = strings.Join(input, "")

	return output
}
