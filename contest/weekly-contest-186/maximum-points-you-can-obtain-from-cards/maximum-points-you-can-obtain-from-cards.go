package leetcode


func maxScore(cardPoints []int, k int) int {
	result := 0
	if len(cardPoints) == k {
		for _ , value := range cardPoints {
			result += value
		}
	} else {
		recursion(cardPoints, 0, &result, 0, k)
	}
	return result
}

func recursion(cardPoints []int, current int, result *int, count int, k int) {

	if count == k {
		*result = max(*result , current)
		return
	}

	if count > k {
		return
	}
	recursion(cardPoints[1:], current + cardPoints[0], result, count + 1, k )
	recursion(cardPoints[0:len(cardPoints) - 1], current + cardPoints[len(cardPoints) - 1], result, count + 1, k )
}

func max(a, b int) int {
	if a < b {
		return  b
	}
	return a
}




