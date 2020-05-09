package main

import (
	"fmt"
)

func main() {
	a := []int{9,0,3,5,7}
	fmt.Println(subsets(a))
}

func subsets(nums []int) [][]int {
	var result [][]int
	recursion(nums, 0, len(nums), []int{}, &result)
	return result
}

func recursion(nums []int, current, max int, tmp []int, result *[][]int) {
	if current >= max {
		t := make([]int , len(tmp))
		copy(t, tmp)
		*result = append(*result, t)
		return
	}

	recursion(nums, current + 1, max, append(tmp, nums[current]) , result)
	recursion(nums, current + 1, max, tmp, result)
}
