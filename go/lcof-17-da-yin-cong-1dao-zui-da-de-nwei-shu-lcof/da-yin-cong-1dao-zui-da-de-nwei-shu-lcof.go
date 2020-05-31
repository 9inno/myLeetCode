package leetcode

import "math"

func printNumbers(n int) []int {
	max := int(math.Pow10(n))
	result := []int{}
	for i := 1; i < max ;i ++ {
		result = append(result, i)
	}
	return result
}
