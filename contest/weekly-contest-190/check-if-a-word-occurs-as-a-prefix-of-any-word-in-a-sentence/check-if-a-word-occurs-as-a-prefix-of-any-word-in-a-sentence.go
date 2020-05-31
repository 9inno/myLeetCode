package leetcode

import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	tmp := strings.Fields(strings.TrimSpace(strings.ToLower(sentence)))

	outLoop : for key , value := range tmp {
		if len(searchWord) > len(value) {
			continue
		}
		for i := 0; i < len(searchWord); i ++ {
			if searchWord[i] != value[i] {
				continue outLoop
			}
		}
		return key + 1
	}
	return -1
}
