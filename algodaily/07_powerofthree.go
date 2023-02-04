package algodaily

import "fmt"

// not working yet
func IsPowerofThree(n int) bool {
	if n/3 == 1 {
		return true
	}
	for n > 0 {
		if n/3 == 1 {
			return true
		}

	}

	// if n <= 3 {
	// 	return false
	// }
	// for n >= 3 {
	// 	if n%3 == 0 {
	// 		n = n / 3
	// 	} else {
	// 		return false
	// 	}
	// }
	return true
}

func CheckIsPowerofThree() {
	fmt.Println("7 : ", IsPowerofThree(7))
	fmt.Println("9 : ", IsPowerofThree(9))
	fmt.Println("1 : ", IsPowerofThree(1))
	fmt.Println("0 : ", IsPowerofThree(0))
}
