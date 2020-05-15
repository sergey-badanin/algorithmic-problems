package backspacecompare

//Solution for the problem:
//Given two strings S and T, return if they are equal considering all backspaces are processed (# means a backspace character).
//
//O(N) for time and O(1) for memory solution
func backspaceCompare(S string, T string) bool {
	condS := condense(S)
	condT := condense(T)
	return condS == condT
}

func condense(s string) string {
	letters := []byte(s)
	shift := 0
	for i := 0; i < len(letters); i++ {
		if letters[i] == '#' {
			if shift > 0 {
				shift--
			}
			continue
		}
		letters[shift] = letters[i]
		shift++
	}
	return string(letters[0:shift])
}
