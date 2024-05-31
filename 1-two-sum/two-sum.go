func twoSum(nums []int, target int) []int {
	m := make(map[int][]int, len(nums))
	for i, n := range nums {
		m[n] = append(m[n], i)
	}
	for k, v := range m {
		if target == 2*k && len(v) > 1 { //in case if we have two repeating ints that sum up to target
			return []int{v[0], v[1]}
		}
		if val, ok := m[target-k]; ok && target != 2*k { // so that we don't return the non-repeating value where k + k == target
			return []int{val[0], v[0]}
		}
	}
	return nil
}