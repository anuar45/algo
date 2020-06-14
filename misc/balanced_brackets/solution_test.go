package solution

import "testing"

func TestBalanced(t *testing.T) {
	tc := []struct {
		input string
		want  bool
	}{
		{"(()())", true},
		{"()()())", false},
	}

	for _, tt := range tc {
		got := CheckBalanced(tt.input)
		if got != tt.want {
			t.Errorf("got: %v, want: %v", got, tt.want)
		}
	}

}
