import (
  "fmt"
    "strings"
)

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}

func equalSubstring(s string, t string, maxCost int) int {

	theLen := len(s)
	prices := make([]int, 0, theLen)

	for i, char := range s {
		prices = append(prices, abs(int(rune(t[i])-char)))
	}

	maxLen := 0
	cost := 0
	curLen := 0

	for i, _ := range prices {
		if maxLen >= theLen-i {
			return maxLen
		}
		curLen = 0
		cost = 0
		for x := i; x < theLen; x++ {
			if cost <= maxCost {
				cost += prices[x]
			}
			if cost <= maxCost {
				curLen++
				if maxLen < curLen {
					maxLen = curLen
				}
				continue
			}
			break
		}
	}

	return maxLen
}