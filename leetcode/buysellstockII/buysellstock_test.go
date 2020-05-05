package buysellstock

import "testing"

func TestMaxProfit(t *testing.T) {
	want := 0
	if got := maxProfit([]int{2, 2, 1}); got != want {
		t.Errorf("For [2, 2, 1] got %v, want %v", got, want)
	}

	want = 0
	if got := maxProfit([]int{2}); got != want {
		t.Errorf("For [2] got %v, want %v", got, want)
	}

	want = 7
	if got := maxProfit([]int{7, 1, 5, 3, 6, 4}); got != want {
		t.Errorf("For [7, 1, 5, 3, 6, 4] got %v, want %v", got, want)
	}

	want = 4
	if got := maxProfit([]int{1, 2, 3, 4, 5}); got != want {
		t.Errorf("For [1, 2, 3, 4, 5] got %v, want %v", got, want)
	}

	want = 0
	if got := maxProfit([]int{7, 6, 4, 3, 1}); got != want {
		t.Errorf("For [7, 6, 4, 3, 1] got %v, want %v", got, want)
	}
}
