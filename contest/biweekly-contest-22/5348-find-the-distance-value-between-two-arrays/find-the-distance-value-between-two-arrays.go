package leetcode

import "math"

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	result := 0
	outLoop: for _ , value1 := range arr1 {
		for _ , value2 := range arr2{
			if abs :=int(math.Abs(float64(value1 - value2))); abs <= d {
				continue outLoop
			}

		}
		result ++
	}
	return result
}