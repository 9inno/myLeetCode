package leetcode

import "sort"

func groupAnagrams(strs []string) [][]string {
	hashMap := map[string][]string{}
	result := [][]string{}
	for _ , value := range strs {
		tmp := []byte(value)
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] < tmp[j]
		})
		hashMap[string(tmp)] = append(hashMap[string(tmp)], value)
	}

	for _ , value := range hashMap {
		result = append(result, value)
	}
	return result
}
