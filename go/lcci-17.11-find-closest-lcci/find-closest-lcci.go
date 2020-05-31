package leetcode

import "math"

func FindClosest(words []string, word1 string, word2 string) int {
	if len(words) == 0 {
		return 0
	}
	p1 := -1
	p2 := -1
	result := math.MaxInt32
	for key , value := range words {
		if value == word1 {
			if p2 != -1 {
				result = min(result , key - p2)
			}
			p1 = key
		}

		if value == word2 {
			if p1 != -1 {
				result = min(result, key - p1)
			}
			p2 = key
		}
	}
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
