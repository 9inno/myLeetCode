package leetcode

func Massage(nums []int) int {

	pre, cur, n := 0, 0, len(nums)
	for i := 0; i < n; i++ {
		pre, cur = cur, max(cur, pre+nums[i])
	}
	return cur
}

func max(a , b int) int {
	if a > b {
		return a
	}
	return b
}