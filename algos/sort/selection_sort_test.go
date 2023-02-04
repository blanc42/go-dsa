package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	A := []int{-1, 0, 11, 6, 2, 12}
	fmt.Print(A)
	SelectionSort(A)
	fmt.Print(" -> ", A)
	want := []int{-1, 0, 2, 6, 11, 12}
	if !reflect.DeepEqual(A, want) {
		t.Errorf("got %v, wanted %v", A, want)
	}
}
