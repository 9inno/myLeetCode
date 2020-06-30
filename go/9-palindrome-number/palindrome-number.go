package leetcode

import "strconv"

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	tmp := x
	newX := 0
	for tmp != 0 {
		newX = tmp % 10 + newX * 10
		tmp /= 10
	}
	return newX == x
}


func IsPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)
	j := len(str) - 1
	for i := 0 ; i < j ;i ++ {
		if str[i] != str[j] {
			return false
		}
		j --
	}
	return true
}