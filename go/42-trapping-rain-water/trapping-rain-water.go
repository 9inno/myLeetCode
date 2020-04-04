package leetcode



func Trap(height []int) int {
	length := len(height)
	if length == 0 {
		return 0
	}
	maxValue := 0
	maxP := 0
	for i := 0; i < length; i ++ {
		if height[i] > maxValue {
			maxValue = height[i]
			maxP = i
		}
	}

	result := 0
	current := height[0]
	for i := 1; i < maxP; i ++ {
		if height[i] < current {
			result += current - height[i]
		}else {
			current = height[i]
		}
	}
	current = height[length - 1]
	for i := length - 2; i > maxP; i -- {
		if height[i] <current {
			result += current - height[i]
		}else {
			current = height[i]
		}
	}
	return result
}