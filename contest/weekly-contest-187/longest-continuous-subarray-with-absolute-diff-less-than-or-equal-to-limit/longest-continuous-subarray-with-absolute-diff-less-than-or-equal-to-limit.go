package main

import "sort"

func main() {
	a := []int{4,2,2,2,4,4,2,2}
	longestSubarray(a ,0)
}

func longestSubarray(nums []int, limit int) int {
	max := len(nums)
	outLoop : for true {
		for i := 0; i + max <= len(nums); i ++  {
			tmp := make([]int, max)
			copy(tmp, nums[i : i + max])
			sort.Ints(tmp)
			if tmp[len(tmp) - 1] - tmp[0] <= limit {
				break outLoop
			}
		}
		max--
	}
	return max
}