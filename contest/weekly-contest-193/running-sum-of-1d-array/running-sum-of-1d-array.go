package leetcode

func runningSum(nums []int) []int {
	result := make([]int , len(nums))
	if len(nums) == 0 {
		return result
	}
	result[0] = nums[0]

	for i := 1 ;i < len(nums); i ++ {
		result[i] = nums[i] + result[i - 1]
	}
	return result
}
