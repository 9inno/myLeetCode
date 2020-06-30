package leetcode

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	i := 0
	j := len(s) - 1
	for i < j {
		if !(s[i] >= 48 && s[i] <=57) &&  !(s[i] >=97 && s[i] <= 122) {
			i ++
			continue
		}

		if !(s[j] >= 48 && s[j] <=57) &&  !(s[j] >=97 && s[j] <= 122) {
			j ++
			continue
		}

		if s[i] != s[j] {
			return false
		}
		i ++
		j --
	}
	return true
}


func isPalindrome1(s string) bool {
	s = strings.ToLower(s)
	result := ""
	reserve := ""
	for _ , value := range s{
		if (value >= 48 && value <=57) ||  (value >=97 && value <= 122) {
			result += string(value)
			reserve = string(value) + reserve
		}
	}
	return result == reserve
}