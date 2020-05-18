package leetcode

import "math"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return math.MinInt32
	}
	result := math.MinInt32
	tmp := 0
	for _ , value :=range nums {
		tmp = max(tmp + value, value)
		result = max(result, tmp)
	}
	return result
}

func max(a , b int) int {
	if a < b {
		return b
	}
	return a
}
