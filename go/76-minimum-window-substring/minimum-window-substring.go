package leetcode

import "math"

func minWindow(s string, t string) string {
	origin, current := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		origin[t[i]]++
	}

	sLen := len(s)
	maxlength := math.MaxInt32
	start, end := -1, -1

	check := func() bool {
		for k, v := range origin {
			if current[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && origin[s[r]] > 0 {
			current[s[r]]++
		}
		for check() && l <= r {
			if r - l + 1 < maxlength {
				maxlength = r - l + 1
				start, end = l, l +maxlength
			}
			if _, ok := origin[s[l]]; ok {
				current[s[l]] -= 1
			}
			l++
		}
	}
	if start == -1 {
		return ""
	}
	return s[start:end]
}
