package solution

import "testing"

func TestSingleNumber(t *testing.T) {
	tc := []struct {
		name  string
		input []int
		want  int
	}{
		{"Test1", []int{2, 2, 1}, 1},
		{"Test2", []int{4, 1, 2, 1, 2}, 4},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			got := singleNumber(tt.input)
			if got != tt.want {
				t.Errorf("want: %v, got: %v", tt.want, got)
			}
		})
	}
}
