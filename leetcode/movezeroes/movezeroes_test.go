package movezeroes

import (
	"testing"
)

func TestSingleNumber(t *testing.T) {
	want := []int{1, 2, 3}
	got := []int{1, 2, 3}
	moveZeroes(got)
	if !equals(got, want) {
		t.Errorf("For [1, 2, 3] got %v, want %v", got, want)
	}

	want = []int{2, 3, 0}
	got = []int{0, 2, 3}
	moveZeroes(got)
	if !equals(got, want) {
		t.Errorf("For [0, 2, 3] got %v, want %v", got, want)
	}

	want = []int{1, 3, 0}
	got = []int{1, 0, 3}
	moveZeroes(got)
	if !equals(got, want) {
		t.Errorf("For [1, 0, 3] got %v, want %v", got, want)
	}

	want = []int{1, 2, 0}
	got = []int{1, 2, 0}
	moveZeroes(got)
	if !equals(got, want) {
		t.Errorf("For [1, 2, 0] got %v, want %v", got, want)
	}

	want = []int{0, 0, 0}
	got = []int{0, 0, 0}
	moveZeroes(got)
	if !equals(got, want) {
		t.Errorf("For [0, 0, 0] got %v, want %v", got, want)
	}

	want = []int{0}
	got = []int{0}
	moveZeroes(got)
	if !equals(got, want) {
		t.Errorf("For [0] got %v, want %v", got, want)
	}

	want = []int{1}
	got = []int{1}
	moveZeroes(got)
	if !equals(got, want) {
		t.Errorf("For [1] got %v, want %v", got, want)
	}

	want = []int{1, 3, 12, 0, 0}
	got = []int{0, 1, 0, 3, 12}
	moveZeroes(got)
	if !equals(got, want) {
		t.Errorf("For [0, 1, 0, 3, 12] got %v, want %v", got, want)
	}
}

func equals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
