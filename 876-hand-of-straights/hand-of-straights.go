func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

	sort.Ints(hand)

	cards := make(map[int]int)
	for _, n := range hand {
		cards[n]++
	}

	for _, card := range hand {
		if cards[card] == 0 {
			continue
		}

		for i := 0; i < groupSize; i++ {
			if cards[card+i] == 0 {
				return false
			}
			cards[card+i]--
		}
	}

	return true
}