package goconcepts

import "testing"

func TestAdd(t *testing.T) {
	// This is an anon function
	Add := func(a, b int) int {
		return a + b
	}
	got := Add(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
