package leetcode

import "math"

func MySqrt(x int) int {

	start := 0
	end := x
	result := 0

	for start <= end {
		mid := start + (end - start) / 2
		if mid * mid <= x {
			result = max(mid , result)
			start = mid + 1
		} else {
			end = mid - 1
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

func MySqrt1(x int) int {
	if x == 0 {
		return 0
	}
	var result float64
	var last float64
	last = 0
	result = 1
	for math.Abs(last - result) > 1e-7 {
		last = result
		result = (result + float64(x) / result) / 2
	}
	return int(result)
}