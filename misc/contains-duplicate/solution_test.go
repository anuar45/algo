package solution

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tc := []struct {
		input []int
		want  bool
	}{
		{[]int{1, 2, 3, 4, 5, 1}, true},
		{[]int{1, 2, 3, 4, 5, 6}, false},
	}

	for _, tt := range tc {
		got := containsDuplicate(tt.input)
		if tt.want != got {
			t.Errorf("want: %v\ngot: %v\n", tt.want, got)
		}
	}
}
