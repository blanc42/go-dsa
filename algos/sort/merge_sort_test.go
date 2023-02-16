package sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	got := MergeSort([]int{2, 1, 3, 4, 5, 9, 7, 6, 8})
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)

	}
}
