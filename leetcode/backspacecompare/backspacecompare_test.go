package backspacecompare

import "testing"

func TestBackspaceCompare(t *testing.T) {
	givenS := "ab#c"
	givenT := "ad#c"
	want := true
	if got := backspaceCompare(givenS, givenT); got != want {
		t.Errorf("For S: %v and T: %v, expected %v, got %v", givenS, givenT, want, got)
	}

	givenS = "ab##"
	givenT = "c#d#"
	want = true
	if got := backspaceCompare(givenS, givenT); got != want {
		t.Errorf("For S: %v and T: %v, expected %v, got %v", givenS, givenT, want, got)
	}

	givenS = "a##c"
	givenT = "#a#c"
	want = true
	if got := backspaceCompare(givenS, givenT); got != want {
		t.Errorf("For S: %v and T: %v, expected %v, got %v", givenS, givenT, want, got)
	}

	givenS = "a#c"
	givenT = "b"
	want = false
	if got := backspaceCompare(givenS, givenT); got != want {
		t.Errorf("For S: %v and T: %v, expected %v, got %v", givenS, givenT, want, got)
	}

	givenS = "abc#d##ef#"
	givenT = "ae"
	want = true
	if got := backspaceCompare(givenS, givenT); got != want {
		t.Errorf("For S: %v and T: %v, expected %v, got %v", givenS, givenT, want, got)
	}
}

func TestCondense(t *testing.T) {
	given := "ab#c"
	want := "ac"
	if got := condense(given); got != want {
		t.Errorf("For %v expected %v, got %v", given, want, got)
	}

	given = "#"
	want = ""
	if got := condense(given); got != want {
		t.Errorf("For %v expected %v, got %v", given, want, got)
	}

	given = "##"
	want = ""
	if got := condense(given); got != want {
		t.Errorf("For %v expected %v, got %v", given, want, got)
	}

	given = "##a"
	want = "a"
	if got := condense(given); got != want {
		t.Errorf("For %v expected %v, got %v", given, want, got)
	}

	given = "abc#d##ef#"
	want = "ae"
	if got := condense(given); got != want {
		t.Errorf("For %v expected %v, got %v", given, want, got)
	}
}
