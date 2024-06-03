func findFromIndex(s string, start int, end int, r rune) (bool, int) {
	for i := start; i <= end; i++ {
		if rune(s[i]) == r {
			return true, i + 1
		}
	}
	return false, -1
}

func appendCharacters(s string, t string) int {
	leftS := 0
	rightS := len(s) - 1
	charsLeft := len(t)

	for _, c := range t {
		if found, j := findFromIndex(s, leftS, rightS, c); found {
			leftS = j
			charsLeft--
			continue
		}
		return charsLeft
	}
	return charsLeft
}