func singleNumber(nums []int) []int {
	m := make(map[int]bool)
	result := make([]int, 0, 2)
	for _, n := range nums {
		if _, ok := m[n]; ok {
			delete(m, n)
			continue
		}
		m[n] = true
	}
	for k, _ := range m {
		result = append(result, k)
	}
	return result
}