func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func scoreOfString(s string) int {
	sum := 0

	for i := 0; i < len(s)-1; i++ {
		sum += int(math.Abs(float64(rune(s[i]) - rune(s[i+1]))))
	}
	return sum
}