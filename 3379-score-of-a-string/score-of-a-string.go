func abs(n rune) rune {
	if n > 0 {
		return n
	}
	return -n
}

func scoreOfString(s string) int {
	var sum rune

	for i := 0; i < len(s)-1; i++ {
		sum += abs(rune(s[i]) - rune(s[i+1]))
	}
	return int(sum)
}