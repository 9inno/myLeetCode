package leetcode

import "strings"

func stringMatching(words []string) []string {
	var result []string
	hashMap := make(map[string]bool)
	for i := 0; i < len(words); i ++ {
		for j := 0; j < len(words) ; j ++ {
			if i == j {
				continue
			}

			if strings.LastIndex(words[i], words[j]) != -1 {
				if _,ok := hashMap[words[j]] ;ok {
					continue
				} else  {
					result = append(result, words[j])
					hashMap[words[j]] = true
				}
			}
		}
	}
	return result
}
