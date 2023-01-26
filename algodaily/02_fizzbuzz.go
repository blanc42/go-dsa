package algodaily

import "fmt"

func Fizzbuzz(n int) {
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Print("fizz")
		} else if i%5 == 0 {
			fmt.Print("buzz")
		} else {
			fmt.Print(i)
		}
	}
	fmt.Print("\n")
}

func CheckFizzbuzz() {
	Fizzbuzz(15)
}
