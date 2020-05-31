package leetcode


func MissingNumber(nums []int) int {
	result := 0
	i := 0
	for ; i < len(nums) ; i ++ {
		result ^= nums[i]
		result ^= i
	}
	return result ^ i
}
