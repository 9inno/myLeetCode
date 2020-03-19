package leetcode


func LongestPalindrome(s string) int {
	charMap := make(map[string]int)
	result := 0
	odd := 0
	for _ , letter := range s {
		if _ , ok := charMap[string(letter)] ; ok {
			charMap[string(letter)] ++
		} else {
			charMap[string(letter)] = 1
		}
	}

	for _ , count := range charMap {
		result += count
		if count % 2 == 1 {
			odd++
		}
	}
	if odd != 0 {
		result = result - odd + 1
	}

	return  result
}