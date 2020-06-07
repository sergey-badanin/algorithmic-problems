package restrictedcalculator

import (
	"fmt"
	"testing"
)

func TestIsClosingCombination(t *testing.T) {
	// comb := []byte{'A', 'B', 'C'}
	// abc := []byte{'A', 'B', 'C'}
	// if !isClosingCombination(comb, abc) {
	// 	t.Errorf("Comb: %v is closing for abc: %v", comb, abc)
	// }

	// comb = []byte{'A', 'B', 'B'}
	// if isClosingCombination(comb, abc) {
	// 	t.Errorf("Comb: %v is not closing for abc: %v", comb, abc)
	// }
}

func TestAddOne(t *testing.T) {
	given := [3]byte{2, 2, 1}
	expect := [3]byte{0, 0, 2}
	addOne(given[:], 2)
	if expect != given {
		t.Errorf("Expected %v but got %v", expect, given)
	}

	given = [3]byte{1, 2, 1}
	expect = [3]byte{2, 2, 1}
	addOne(given[:], 2)
	if expect != given {
		t.Errorf("Expected %v but got %v", expect, given)
	}

	given = [3]byte{1, 0, 0}
	expect = [3]byte{2, 0, 0}
	addOne(given[:], 2)
	if expect != given {
		t.Errorf("Expected %v but got %v", expect, given)
	}

	given = [3]byte{2, 2, 0}
	expect = [3]byte{0, 0, 1}
	addOne(given[:], 2)
	if expect != given {
		t.Errorf("Expected %v but got %v", expect, given)
	}
}

func TestIterateCombinations(t *testing.T) {

	result := iterateCombinations(3, 2, func(comb []byte) bool {
		if comb[0] == 0 && comb[1] == 0 && comb[2] == 2 {
			return true
		}
		return false
	})

	if !(result[0] == 0 && result[1] == 0 && result[2] == 2) {
		t.Errorf("Wrong reslut %v", result)
	}
}

func TestConformsAB(t *testing.T) {
	if !conformsAB([]byte{1, 1, 0, 1}, 22) {
		t.Error("Should conform")
	}
}

func TestConformsABC(t *testing.T) {
	if !conformsABC([]byte{0, 1, 2, 1}, 26) {
		t.Error("Should conform")
	}
}

func TestSearchConformingAB(t *testing.T) {
	result := searchConformingAB()
	fmt.Println(result)
}

func TestSearchConformingABC(t *testing.T) {
	result := searchConformingABC()
	fmt.Println(result)
}
