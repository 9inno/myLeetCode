package leetcode



func MaxArea(height []int) int {
	i := 0
	j := len(height) - 1
	result := 0
	for i < j {
		result = max(result, (j - i) * min(height[i], height[j]))
		if height[i] > height[j] {
			j --
		} else {
			i ++
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

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}