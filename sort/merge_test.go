package sort

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tc1 := []int{1, 3, 5, 7, 9}
	tc2 := []int{2, 4, 6, 8}

	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	got := Merge(tc1, tc2)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("should be equal: want: %v, got: %v", want, got)
	}
}

func TestMergeSort(t *testing.T) {
	tc1 := []int{8, 4, 9, 6, 2, 7, 3, 5, 1}

	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	got := MergeSort(tc1)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("should be equal: want: %v, got: %v", want, got)
	}
}
