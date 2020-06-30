package leetcode

func longestSubarray(nums []int) int {
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	leftCount := 0
	for key , value := range nums {
		if value == 0 {
			left[key] = leftCount
			leftCount = 0
		} else {
			leftCount ++
		}
	}
	rightCount := 0
	for i := len(nums) - 1; i >=0 ; i -- {
		if nums[i] == 0 {
			right[i] = rightCount
			rightCount = 0
		}else {
			rightCount ++
		}
	}

	max := 0

	for i:= 0; i < len(nums); i ++ {
		if left[i] + right[i] > max {
			max = left[i] + right[i]
		}
	}
	if max == 0 && leftCount == rightCount && leftCount != 0{
		return leftCount - 1
	}
	return max
}
