package sequentialpages

import (
	"fmt"
	"testing"
)

func TestDigitsNumberToPage(t *testing.T) {
	page := digitsNumberToPage(990)
	fmt.Println(page)
}

func TestMaxEncountersOfDigit(t *testing.T) {
	got := maxEncountersOfDigit('1', 1)
	expect := 9
	if got != expect {
		t.Errorf("Got %v expect %v", got, expect)
	}

	got = maxEncountersOfDigit('1', 2)
	expect = 10
	if got != expect {
		t.Errorf("Got %v expect %v", got, expect)
	}

	got = maxEncountersOfDigit('1', 4)
	expect = 11
	if got != expect {
		t.Errorf("Got %v expect %v", got, expect)
	}

	got = maxEncountersOfDigit('1', 261)
	expect = 809
	if got != expect {
		t.Errorf("Got %v expect %v", got, expect)
	}

	got = maxEncountersOfDigit('1', 260)
	fmt.Println(got)
}
