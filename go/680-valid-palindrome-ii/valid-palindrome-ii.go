package leetcode


func validPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i <= j {
		if s[i] == s[j] {
			i ++
			j --
			continue
		}
		flag1 := true
		flag2 := true

		for x , y := i + 1, j ;x <= y; x , y = x + 1, y - 1 {
			if s[x] != s[y] {
				flag1 = false
				break
			}
		}
		for x , y := i , j - 1 ;x <= y; x , y = x + 1, y - 1 {
			if s[x] != s[y] {
				flag2 = false
				break
			}
		}
		return flag1 || flag2
	}
	return true
}
