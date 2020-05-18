package leetcode


func searchMatrix(matrix [][]int, target int) bool {
	result := false
	for _ , value := range matrix {
		if len(value) == 0 {
			continue
		}
		if value[0] <= target && value[len(value) - 1] >= target {
			result = result || dichotomy(value, target, 0 , len(value) - 1)
		}
	}
	return result
}

func dichotomy(nums []int , target, start , end int) bool {
	if start > end {
		return false
	}
	mid := start + (end -start) / 2
	if nums[start] == target {
		return true
	}
	if nums[end] == target {
		return true
	}
	if nums[mid] == target {
		return true
	}

	if target < nums[mid] {
		return dichotomy(nums, target, start + 1, mid - 1)
	} else {
		return dichotomy(nums, target, mid + 1 , end - 1)
	}
}