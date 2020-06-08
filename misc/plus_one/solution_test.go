package solution

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	tc := []struct {
		input []int
		want  []int
	}{
		{
			input: []int{4, 3, 2, 1},
			want:  []int{4, 3, 2, 2},
		},
		{
			input: []int{4, 3, 9, 9},
			want:  []int{4, 4, 0, 0},
		},
		{
			input: []int{9, 9, 9},
			want:  []int{1, 0, 0, 0},
		},
	}

	for _, tt := range tc {
		got := plusOne(tt.input)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("error want: %v, got: %v", tt.want, got)
		}
	}
	// t.Fatal("not implemented")
}
