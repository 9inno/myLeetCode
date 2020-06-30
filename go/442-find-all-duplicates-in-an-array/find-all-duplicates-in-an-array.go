package leetcode

import "math"

func findDuplicates(nums []int) []int {
	result := []int{}
	for _ , value := range nums {
		if nums[int(math.Abs(float64(value))) - 1] < 0 {
			result = append(result, int(math.Abs(float64(value))))
		} else {
			nums[int(math.Abs(float64(value))) - 1] *= -1
		}
	}
	return result
}
