package algodaily

import (
	"fmt"
	"regexp"
	"strings"
)

func ReverseOnlyAlphabetical(str string) {
	input := strings.Split(str, "")
	l := len(input)
	re := regexp.MustCompile(`[a-z]|[A-Z]`)
	j := l - 1
	for i := 0; i < j; i++ {
		if re.Match([]byte(input[i])) {
			for ; i < j; j-- {
				if re.Match([]byte(input[j])) {
					temp := input[i]
					input[i] = input[j]
					input[j] = temp
					j--
					break
				}
			}
		}
	}
	out := strings.Join(input, "")
	fmt.Println(out)
}

func CheckReverseOnlyAlphabetical() {
	ReverseOnlyAlphabetical("sea!$hells3")
}
