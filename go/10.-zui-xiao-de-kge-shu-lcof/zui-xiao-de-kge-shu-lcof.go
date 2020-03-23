package leetcode

import "sort"

func GetLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[: k]
}
