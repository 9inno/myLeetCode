package leetcode


func LengthOfLongestSubstring(s string) int {
	result := 0
	hashMap := map[string]bool{}

	i := 0

	for j := 0; j < len(s) ; j ++ {
		if _ , ok := hashMap[string(s[j])] ; !ok {
			hashMap[string(s[j])] = true
		} else {
			result = max(result, len(hashMap))
			for s[i] != s[j]{
				delete(hashMap, string(s[i]))
				i ++
			}
			i ++
		}
	}
	result = max(result, len(hashMap))

	return result
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}