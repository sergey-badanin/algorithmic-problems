package maximumsubarray

import "testing"

func TestMaxSubArray(t *testing.T) {
	want := 6
	if got := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}); got != want {
		t.Errorf("For [-2, 1, -3, 4, -1, 2, 1, -5, 4] got %v, want %v", got, want)
	}
}
