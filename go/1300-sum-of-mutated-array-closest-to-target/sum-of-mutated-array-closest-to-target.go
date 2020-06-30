package leetcode

import "math"

func findBestValue(arr []int, target int) int {
	length := len(arr)
	avg := target / length
	pre := math.MaxFloat32

	for i := avg ; ; i ++ {
		sum := 0
		for _ , value := range  arr {
			sum += min(value, i)
		}
		if math.Abs(float64(sum - target)) >= pre {
			return i - 1
		}
		pre = math.Abs(float64(sum - target))
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}