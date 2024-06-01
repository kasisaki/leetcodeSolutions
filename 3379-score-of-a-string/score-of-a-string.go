func scoreOfString(s string) int {
	sum := 0
	diff := 0

	for i := 0; i < len(s)-1; i++ {
		a := rune(s[i])
		b := rune(s[i+1])
		if a != b {
			diff = int(math.Abs(float64(a - b)))
			sum += diff
		}
	}
	return sum
}