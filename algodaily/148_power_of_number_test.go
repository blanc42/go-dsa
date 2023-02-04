package algodaily

import (
	"testing"
)

func TestPowerOfNumber(t *testing.T) {
	got := PowerOfNumber(1, 0)
	var want float64 = 1

	if got != want {
		t.Errorf("got %f, wanted %f", got, want)

	}
	got = PowerOfNumber(2, 5)
	want = 32

	if got != want {
		t.Errorf("got %f, wanted %f", got, want)

	}
}
