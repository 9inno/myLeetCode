package leetcode

import "sort"

func average(salary []int) float64 {
	length := len(salary)
	sort.Ints(salary)
	result := 0.0
	count := 0
	for i := 1; i < length - 1; i ++ {
		result += float64(salary[i])
		count ++
	}
	return result / float64(count)
}