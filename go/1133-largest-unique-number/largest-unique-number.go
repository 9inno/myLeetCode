package leetcode

func LargestUniqueNumber(A []int) int {
	hashMap := map[int]int{}
	for _ , value := range A {
		hashMap[value] ++
	}
	result := -1
	for key , count := range hashMap {
		if count == 1 {
			result = max(result , key)
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
