package leetcode

func kidsWithCandies(candies []int, extraCandies int) []bool {
	tmp := 0
	for _, value := range candies {
		tmp = max(tmp, value)
	}
	var result []bool
	for _ , value := range candies {
		if value + extraCandies >= tmp {
			result = append(result, true)
		} else {
			result = append(result, false)
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
