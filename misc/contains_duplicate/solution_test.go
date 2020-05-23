package solution

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tc := []struct {
		name  string
		input []int
		want  bool
	}{
		{"trueCase", []int{1, 2, 3, 4, 5, 1}, true},
		{"falseCase", []int{1, 2, 3, 4, 5, 6}, false},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			got := containsDuplicate(tt.input)
			if tt.want != got {
				t.Errorf("want: %v\ngot: %v\n", tt.want, got)
			}
		})
	}
}

func BenchmarkContainsDuplicate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		containsDuplicate([]int{1, 2, 3, 4, 5, 6, 8, 9})
	}
}
