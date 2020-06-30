package leetcode

func shuffle(nums []int, n int) []int {
	result := []int{}
	for i := 0 ; i < len(nums) - n ; i ++ {
		result = append(result, nums[i])
		result = append(result, nums[i + n ])
	}
	return result
}
