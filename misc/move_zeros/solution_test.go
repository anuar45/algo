package solution

import "reflect"
import "testing"

func TestSolution(t *testing.T) {
	input := []int{0, 1, 6, 4, 0, 8, 5, 0, 4}
	want := []int{1, 6, 4, 8, 5, 4, 0, 0, 0}

	moveZeros(input)

	if !reflect.DeepEqual(want, input) {
		t.Errorf("Error: %v != %v", want, input)
	}
}
