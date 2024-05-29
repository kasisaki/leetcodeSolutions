func numSteps(s string) int {
    steps := 0
	i := len(s) - 1
    
	for i >= 1 {
		if s[i]-'0' == 0 {
			steps++
			s = s[0:i]
		} else {
			steps++
			s = addBinary(s)
		}
		i = len(s) - 1
	}

	return steps
}

func addBinary(s string) string {
	result := ""
	carry := 1

	for i := len(s) - 1; i >= 0; i-- {
		bit := (s[i] - '0') + byte(carry)

		if bit > 1 {
			result = "0" + result
			carry = 1
		} else {
			result = fmt.Sprintf("%d", bit) + result
			carry = 0
		}
	}

	if carry == 1 {
		result = "1" + result
	}

	return result
}