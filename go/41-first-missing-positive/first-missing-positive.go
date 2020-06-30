package leetcode

import "sort"

func firstMissingPositive(nums []int) int {
	sort.Ints(nums)
	i := 1
	for key , value := range nums {
		if value <= 0 {
			continue
		}
		if key < len(nums) - 1 && value == nums[key + 1] {
			continue
		}
		if value != i {
			break
		}
		i ++
	}
	return i
}


func firstMissingPositive1(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i + 1 {
			return i + 1
		}
	}
	return n + 1

}
