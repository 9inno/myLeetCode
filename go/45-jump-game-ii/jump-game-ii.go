package leetcode

func Jump(nums []int) int {

	result := make([]int, len(nums))
	for i := 0 ; i < len(nums); i ++ {
		result[i] = i
	}

	for i := 0 ; i < len(nums); i ++ {
		for j := i ; j < len(nums) && j <= i + nums[i]; j ++ {
			result[j] = min(result[i] + 1, result[j])
		}
	}
	return result[len(nums) - 1]
}

func min(a , b int) int {
	if a > b {
		return b
	}
	return a
}


func Jump1(nums []int) int {
	result := 0
	furthest := 0
	end := 0
	for i := 0; i < len(nums) - 1; i ++ {
		furthest = max(furthest, nums[i] + i)
		if i == end {
			end = furthest
			result ++
		}
	}
	return result
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}