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




func maxScore1(cardPoints []int, k int) int {
	n := len(cardPoints)
	leftSum := 0
	rightSum := 0
	for i := 0; i < k; i ++ {
		leftSum += cardPoints[i]
	}
	result := leftSum
	for j := 1; j <=k ; j ++ {
		leftSum -= cardPoints[k - j]
		rightSum += cardPoints[n -j]
		result = max(result, leftSum +rightSum)
	}
	return result

}