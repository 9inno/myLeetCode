package leetcode

func maxScoreSightseeingPair(A []int) int {
	result := 0
	maxValue := A[0]

	for i := 1; i < len(A); i ++ {
		result = max(result, maxValue + A[i] - i)
		maxValue = max(maxValue , A[i] + i )
	}

	return result
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}