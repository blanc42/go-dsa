package algodaily

import "testing"

func TestLonelyNumber(t *testing.T) {
	LonelyList := []int{2, 1, 1, 4, 3, 4, 2, 13, 13, 14, 14, 16, 17, 17, 16}
	got := LonelyNumber(LonelyList)
	var want int = 3

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)

	}
}
