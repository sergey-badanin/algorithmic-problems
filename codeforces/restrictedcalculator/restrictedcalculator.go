package restrictedcalculator

/*
 * A crazed inventor made a calculator with only two buttons, A and B.
 * It's not very useful, but there's a surprising amount you can do with it;
 * when button A is pressed, the current number on the display is increased by 3,
 * and when the button B is pressed, the current number is multiplied by 2.
 */

func searchConformingAB() []byte {
	for size := 2; size < 5; size++ {
		comb := iterateCombinations(size, 1, func(comb []byte) bool { return conformsAB(comb, 17) })
		if comb != nil {
			return comb
		}
	}
	return nil
}

func searchConformingABC() []byte {
	for size := 2; size < 7; size++ {
		comb := iterateCombinations(size, 2, func(comb []byte) bool { return conformsABC(comb, 47) })
		if comb != nil {
			return comb
		}
	}
	return nil
}

// A increase by 3
// B multiply by 2
func conformsAB(comb []byte, expect int) bool {
	number := 2
	for _, v := range comb {
		if v == 0 {
			number += 3
			continue
		}
		if v == 1 {
			number *= 2
		}
	}
	if number == expect {
		return true
	}
	return false
}

// A increase by 3
// B multiply by 2
// C increase by 1
func conformsABC(comb []byte, expect int) bool {
	number := 3
	for _, v := range comb {
		if v == 0 {
			number += 3
			continue
		}
		if v == 1 {
			number *= 2
		}
		if v == 2 {
			number++
		}
	}
	if number == expect {
		return true
	}
	return false
}

/*
 * Generate all possible combinations
 */
func iterateCombinations(size int, highest byte, conforms func(comb []byte) bool) []byte {
	comb := make([]byte, size)

	for !isClosingCombination(comb, highest) {
		addOne(comb, highest)
		if conforms(comb) {
			return comb
		}
	}
	return nil
}

func addOne(comb []byte, highest byte) {
	incr := 1
	index := 0

	for incr > 0 && index < len(comb) {
		if comb[index] < highest {
			comb[index]++
			incr = 0
			continue
		}
		if comb[index] == highest {
			comb[index] = 0
			index++
		}
	}
}

func isClosingCombination(comb []byte, highest byte) bool {
	for _, v := range comb {
		if v != highest {
			return false
		}
	}
	return true
}
