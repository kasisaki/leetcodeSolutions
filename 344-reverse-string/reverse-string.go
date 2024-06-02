func reverseString(s []byte) {
	var a, b byte
	j := len(s)

	for i := 0; i < len(s) / 2; i++ {
		a = s[i]
		b = s[j-1-i]
		s[i] = b
		s[j-1-i] = a
	}
}