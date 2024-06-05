func commonChars(words []string) []string {
	minLen := 0
	indOfMin := 0

	for i, word := range words {
		lenW := len(word)
		if lenW < minLen || minLen == 0 {
			minLen = lenW
			indOfMin = i
		}
	}

	minWord := words[indOfMin]
	minWordMap := make(map[rune]int)

	for _, c := range minWord {
		minWordMap[c]++
	}

	for i, word := range words {
		if i == indOfMin {
			continue
		}

		wordMap := make(map[rune]int)
		for _, char := range word {
			if minWordMap[char] > 0 {
				wordMap[char]++
			}
		}

		for k, v := range minWordMap {
			if wordMap[k] == 0 {
				delete(minWordMap, k)
				continue
			}
			if wordMap[k] < v {
				minWordMap[k] = wordMap[k]
			}
		}
	}

	result := make([]string, 0)
	for k, v := range minWordMap {
		for i := 0; i < v; i++ {
			result = append(result, string(k))
		}
	}

	return result
}