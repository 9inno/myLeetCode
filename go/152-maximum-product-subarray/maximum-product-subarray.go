package main

func main() {
	nums := []int{-2, 3, -4}
	maxProduct(nums)
}
func maxProduct(nums []int) int {
	result := nums[0]
	maxValue := nums[0]
	minValue := nums[0]

	for i := 1; i < len(nums); i ++ {
		tmpMax := maxValue
		tmpMin := minValue
		maxValue = max(tmpMax * nums[i], max(tmpMin * nums[i],nums[i]))
		minValue = min(tmpMin * nums[i], min(tmpMax * nums[i], nums[i]))
		result = max(maxValue, result)
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