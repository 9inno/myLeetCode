package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	length := len(nums)

	for i := 0; i < length; i ++ {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}
		z := length - 1
		target := nums[i] * -1
		for j := i + 1 ; j < length ; j ++ {
			if j > i + 1  && nums[j] == nums[j - 1] {
				continue
			}

			for j < z && nums[j] + nums[z] > target {
				z --
			}

			if j == z {
				break
			}

			if nums[j] + nums[z] == target {
				result = append(result, []int{nums[i], nums[j], nums[z]})
			}
		}
	}
	return result
}
