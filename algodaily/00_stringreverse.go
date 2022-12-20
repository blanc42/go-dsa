package algodaily

import "strings"

// Reverse a string Algodaily
func StringReverse(i string) (output string) {
	input := strings.Split(i, "")
	j := len(input)
	for i := 0; i < j; i++ {
		temp := input[i]
		input[i] = input[j-1]
		input[j-1] = temp
		j--
	}
	output = strings.Join(input, "")

	return output
}
