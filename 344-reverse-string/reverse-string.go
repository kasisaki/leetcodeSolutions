func reverseString(s []byte) {
	var a, b byte
	l := len(s) / 2
	j := len(s) - 1

	for i := 0; i < l; i++ {
		a = s[i]
		b = s[j-i]
		s[i] = b
		s[j-i] = a
	}
}