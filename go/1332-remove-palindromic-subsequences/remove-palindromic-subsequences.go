package leetcode

import "strings"

func removePalindromeSub(s string) int {
	if len(s) == 0 {
		return 0
	}
	tmp := strings.Split(s, "")
	i := 0
	j := len(tmp) - 1
	for i < j {
		tmp[i], tmp[j] = tmp[j], tmp[i]
		i++
		j--
	}
	if s == strings.Join(tmp, "") {
		return 1
	}
	return 2
}
