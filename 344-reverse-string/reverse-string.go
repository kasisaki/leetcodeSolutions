func reverseString(s []byte) {
	var a byte
	l := len(s) / 2
	j := len(s) - 1

	for i := 0; i < l; i++ {
		a = s[i]
		s[i] = s[j]
		s[j] = a
        j--
	}
}