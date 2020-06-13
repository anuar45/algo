package sort

import (
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	tc := []int{8, 4, 9, 6, 2, 7, 3, 5, 1}

	want := []int{5, 4, 1, 6, 2, 7, 3, 8, 9}

	pivotGot := partitition(tc, 0, len(tc)-1)

	if !reflect.DeepEqual(tc, want) {
		t.Errorf("sort want: %v, got: %v", want, tc)
	}

	pivotWant := 7
	if pivotGot != pivotWant {
		t.Errorf("pivot want: %v, got: %v", pivotWant, pivotGot)
	}

}

func TestQuick(t *testing.T) {
	tc := []int{8, 4, 9, 6, 2, 7, 3, 5, 1}

	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	Quick(tc, 0, len(tc)-1)

	if !reflect.DeepEqual(tc, want) {
		t.Errorf("sort want: %v, got: %v", want, tc)
	}

}
