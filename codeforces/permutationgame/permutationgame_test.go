package permutationgame

import (
	"fmt"
	"testing"
)

func TestGeneratePermutations(t *testing.T) {
	expected := []int{2, 2, 1, 0}
	answer := generatePermutations([]int{0, 1, 2, 3}, func(perm []int) bool {
		if equals(expected, evaluate(perm)) {
			return true
		}
		return false
	})
	fmt.Printf("For answers: %v word is:%v\n", expected, answer)

	expected = []int{1, 2, 0, 0}
	answer = generatePermutations([]int{0, 1, 2, 3}, func(perm []int) bool {
		if equals(expected, evaluate(perm)) {
			return true
		}
		return false
	})
	fmt.Printf("For answers: %v word is:%v\n", expected, answer)

	expected = []int{4, 1, 3, 1, 0, 0}
	answer = generatePermutations([]int{0, 1, 2, 3, 4, 5}, func(perm []int) bool {
		if equals(expected, evaluate(perm)) {
			return true
		}
		return false
	})
	fmt.Printf("For answers: %v word is:%v\n", expected, answer)

	expected = []int{6, 4, 0, 4, 0, 0, 2, 1, 0}
	answer = generatePermutations([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}, func(perm []int) bool {
		if equals(expected, evaluate(perm)) {
			return true
		}
		return false
	})
	fmt.Printf("For answers: %v word is:%v\n", expected, answer)
}

func TestEvaluate(t *testing.T) {
	given := [4]int{3, 2, 0, 1}
	expect := []int{2, 2, 1, 0}
	if got := evaluate(given[:]); !equals(got, expect) {
		t.Errorf("Got: %v, expect %v", got, expect)
	}
}
