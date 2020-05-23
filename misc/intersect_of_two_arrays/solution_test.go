package solution

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	tc := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  []int
	}{
		{"Test1", []int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
		{"Test2", []int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			got := intersect(tt.nums1, tt.nums2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("should be equal got: %v\nwant: %v", got, tt.want)
			}
		})
	}
}
