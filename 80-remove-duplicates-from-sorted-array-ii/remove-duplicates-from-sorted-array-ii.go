package main

import "fmt"

func removeDuplicates(nums []int) int {
	l := len(nums)
	count := 1
	beginning := true
	prev := 0
	moves := 0

	for i, n := range nums {
		if beginning {
			beginning = false
			prev = n
			continue
		}
		if n == prev {

			count++
			if count > 2 {
				nums[i] = n ^ (1 << 14) // move out of the range
				moves++
			}
			continue
		}
		prev = n
		count = 1
	}
	badIndex := -1
	for i, n := range nums {
		if n > 10000 || n < -10000 {
			if badIndex == -1 {
				badIndex = i
			}
			continue
		}
		if badIndex == -1 {
			continue
		}
		nums[badIndex] = n
		badIndex++
	}

	return l - moves
}