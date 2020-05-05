package singlenumber

import "testing"

func TestSingleNumber(t *testing.T) {
	want := 1
	if got := singleNumber([]int{2, 2, 1}); got != want {
		t.Errorf("For [2, 2, 1] got %v, want %v", got, want)
	}

	want = 4
	if got := singleNumber([]int{4, 1, 2, 1, 2}); got != want {
		t.Errorf("For [4, 1, 2, 1, 2] got %v, want %v", got, want)
	}
}
