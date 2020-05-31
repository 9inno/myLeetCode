package leetcode

func Rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	if length == 2 {
		return max(nums[0], nums[1])
	}
	dp := make([]int, len(nums))
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])
	for i := 2 ;i < len(nums); i ++ {
		dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])
	}

	return dp[length - 1]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
