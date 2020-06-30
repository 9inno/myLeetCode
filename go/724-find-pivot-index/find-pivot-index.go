package leetcode

func pivotIndex(nums []int) int {

	left := make([]int, len(nums))
	right := make([]int, len(nums))
	for i := 0 ; i < len(nums); i ++ {
		if i == 0 {
			left[i] = 0
			continue
		}
		if i == 1 {
			left[i] = nums[i - 1]
			continue
		}
		left[i] = left[i - 1] + nums[i - 1]
	}

	for i := len(nums) - 1 ; i >= 0 ; i -- {
		if i == len(nums) - 1 {
			right[i] = 0
			continue
		}
		if i == len(nums) - 2 {
			right[i] = nums[i + 1]
			continue
		}
		right[i] = right[i + 1] + nums[i + 1]
	}

	for i := 0; i < len(nums); i ++ {
		if left[i] == right[i] {
			return i
		}
	}
	return -1
}
