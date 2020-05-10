package rotate

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tc := []struct {
		nums []int
		k    int
		want []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			3,
			[]int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			[]int{1, 2},
			2,
			[]int{1, 2},
		},
		{
			[]int{1, 2},
			0,
			[]int{1, 2},
		},
		{
			[]int{1, 2},
			3,
			[]int{2, 1},
		},
	}

	for _, tt := range tc {
		rotate(tt.nums, tt.k)
		if !reflect.DeepEqual(tt.nums, tt.want) {
			t.Errorf("not equal. want: %v, got: %v", tt.want, tt.nums)
		}
	}
}

func TestReverse(t *testing.T) {
	tc := []struct {
		nums  []int
		start int
		end   int
		want  []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			2,
			5,
			[]int{1, 2, 6, 5, 4, 3, 7},
		},
	}

	for _, tt := range tc {
		reverse(tt.nums, tt.start, tt.end)
		//fmt.Println(tt.nums, tt.want)
		if !reflect.DeepEqual(tt.nums, tt.want) {
			t.Errorf("not equal. want: %v, got: %v", tt.want, tt.nums)
		}
	}

}
