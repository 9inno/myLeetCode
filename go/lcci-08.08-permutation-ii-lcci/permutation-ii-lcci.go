package leetcode



func permutation(S string) []string {
	result := []string{}
	hashMap := map[string]bool{}
	perm([]byte(S), &result, 0, len(S), hashMap)
	return result
}

func perm(arr []byte, result *[]string, n int, count int, hasMap map[string]bool) {
	if count == 0 {
		return
	}
	if count == n {
		if _, ok := hasMap[string(arr)]; !ok {
			hasMap[string(arr)] = true
			*result = append( *result, string(arr))
		}
	}
	for i := n ; i < count; i ++ {
		arr[n], arr[i] = arr[i] , arr[n]
		perm(arr, result, n + 1, count, hasMap)
		arr[n], arr[i] = arr[i] , arr[n]
	}
}
