package leetcode

import (
	"strconv"
	"strings"
)

func hasAllCodes(s string, k int) bool {
	n := 1 << k
	for i := 0 ; i < n ; i ++ {
		tmp := convertToBin(i)
		for len(tmp) < k {
			tmp = "0" + tmp
		}
		if !strings.Contains(s , tmp) {
			return false
		}
	}
	return true
}



func convertToBin(num int) string {
	s := ""
	if num == 0 {
		return "0"
	}
	for ;num > 0 ; num /= 2 {
		lsb := num % 2
		s = strconv.Itoa(lsb) + s
	}
	return s
}

