func reverseString(s []byte) {
	var a, b byte
	l := len(s) / 2
	j := len(s)

	for i := 0; i < l; i++ {
		a = s[i]
		b = s[j-1-i]
		s[i] = b
		s[j-1-i] = a
	}
}