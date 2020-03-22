package leetcode

import "sort"

func MinIncrementForUnique(A []int) int {
	sort.Ints(A)
	result := 0
	for i := 1; i < len(A) ; i ++ {
		if A[i] <= A[i-1] {
			result += A[i-1]-A[i] + 1
			A[i] = A[i-1] + 1
		}
	}
	return result
}
