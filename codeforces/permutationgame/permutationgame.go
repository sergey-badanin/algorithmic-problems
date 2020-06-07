package permutationgame

/*
 * Alice and Bob are playing the following game.
 * Alice secretly writes down the first n letters of the alphabet in some order that she decides.
 * So, if n=3, she can write one of the six words: "ABC", "ACB", "BAC", "BCA", "CAB" or "CBA".
 * For example, consider the situation when n=4 and Alice wrote down the word "DCAB".
 * Then Bob asks Alice several questions, one for each letter from 'A' to 'D':
 * -How many letters, that appear in the alphabet after 'A', go before 'A' in your word? Alice answers 2, because letters 'C' and 'D' stand before 'A'.
 * -How many letters, that appear in the alphabet after 'B', go before 'B' in your word? Alice again answers 2: letters 'C' and 'D' are before 'B' in the word.
 * -How many letters, that appear in the alphabet after 'C', go before 'C' in your word? Alice answers 1: D is before 'C' in the word.
 * -How many letters, that appear in the alphabet after 'D', go before 'D' in your word? Alice answers 0.
 * Lets write down all Alice's answers as a sequence (2,2,1,0).
 * Now Bob should try to guess the word from the information that he got.
 * Can you play this game instead of Bob??
 */

/*
 * Heap's algorithm to generate all permutations
 */
func generatePermutations(arr []int, expectedEvaluation func(perm []int) bool) []int {
	c := make([]int, len(arr))

	if expectedEvaluation(arr) {
		return arr
	}

	i := 0
	for i < len(arr) {
		if c[i] < i {
			if i%2 == 0 {
				swap(arr, 0, i)
			} else {
				swap(arr, c[i], i)
			}
			if expectedEvaluation(arr) {
				return arr
			}
			c[i]++
			i = 0
		} else {
			c[i] = 0
			i++
		}
	}
	return nil
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func evaluate(perm []int) []int {
	answer := make([]int, len(perm))
	highest := len(perm)

	for let := 0; let < highest; let++ {
		for _, v := range perm {
			if v > let {
				answer[let]++
			}
			if v == let {
				break
			}
		}
	}
	return answer
}

func equals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}
