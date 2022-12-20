package utils

func Swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
