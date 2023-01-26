package utils

// takes two pointers to integers and swaps them
func SwapInt(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

// takes two pointer to strings/characters and swaps them
func SwapString(a *string, b *string) {
	temp := *a
	*a = *b
	*b = temp
}
