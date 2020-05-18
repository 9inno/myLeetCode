package leetcode


func maxPower(s string) int {
	if len(s) == 0 {
		return 0
	}
	result := 0
	pre := ""
	tmp := 1
	for _ , letter := range s {
		if string(letter) != pre {
			pre = string(letter)
			result = max(result, tmp)
			tmp = 1
		}else {
			tmp ++
		}
	}
	result = max(result, tmp)
	return result
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
