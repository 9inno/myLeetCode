package leetcode

func maxSubArray(nums []int) int {
	maxValue := nums[0]
	for i := 1; i < len(nums); i ++ {
		nums[i] = max(nums[i], nums[i] + nums[i - 1])
		maxValue = max(maxValue, nums[i])
	}
	return maxValue
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
