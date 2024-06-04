func longestPalindrome(s string) int {
	m := make(map[rune]int)
    var singleLetterPresent bool
	var doubles int

	for _, c := range s {
		m[c]++
	}
	
	for _, v := range m {
		if v%2 == 1 {
			singleLetterPresent = true
		}
		doubles += v / 2
	}

	if singleLetterPresent {
		return doubles*2 + 1
	}
    
	return doubles * 2
}