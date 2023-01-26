package algodaily

import "fmt"

func IsPalindrome(str string) bool {
	// input := strings.Split(str, "")
	l := len(str)
	j := l - 1
	// <= since it has to check for the middle as well
	for i := 0; i <= j; i++ {
		if str[i] == str[j] {
			j--
			continue
		} else {
			return false
		}

	}
	return true
}

func CheckIsPalindrom() {
	fmt.Println(IsPalindrome("racecar"))
}
