package happynumber

import "testing"

func TestIsHappy(t *testing.T) {
	want := true
	if got := isHappy(1); got != want {
		t.Errorf("isHappy(1) = %v, want %v", got, want)
	}

	want = true
	if got := isHappy(19); got != want {
		t.Errorf("isHappy(19) = %v, want %v", got, want)
	}

	want = true
	if got := isHappy(13); got != want {
		t.Errorf("isHappy(13) = %v, want %v", got, want)
	}

	want = false
	if got := isHappy(20); got != want {
		t.Errorf("isHappy(20) = %v, want %v", got, want)
	}
}

func TestSumOfSquaredDigits(t *testing.T) {
	want := 1
	if got := sumOfSquaredDigits(1); got != want {
		t.Errorf("sumSquareDigitd(1) = %v, want %v", got, want)
	}

	want = 91
	if got := sumOfSquaredDigits(139); got != want {
		t.Errorf("sumSquareDigitd(139) = %v, want %v", got, want)
	}
}
