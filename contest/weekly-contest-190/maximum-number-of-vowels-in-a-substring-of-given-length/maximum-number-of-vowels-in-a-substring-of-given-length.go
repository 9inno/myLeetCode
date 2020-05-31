package main

func main() {
	maxVowels("ibpbhixfiouhdljnjfflpapptrxgcomvnb",34)
}


func maxVowels(s string, k int) int {
	if len(s) == 0 {
		return 0
	}
	pre := make([]int, len(s))
	if s[0] == 'a' || s[0] == 'e' ||s[0] == 'i' ||s[0] == 'o' || s[0] == 'u' {
		pre[0] = 1
	}
	for i := 1; i < len(pre); i ++ {
		if s[i] == 'a' || s[i] == 'e' ||s[i] == 'i' ||s[i] == 'o' || s[i] == 'u' {
			pre[i] = pre[i - 1] + 1
		} else {
			pre[i] = pre[i - 1]
		}
	}

	x := 0
	y := x + k - 1
	maxValue := 0
	for y < len(pre) {
		tmpValue := pre[y] - pre[x]
		if s[x] == 'a' || s[x] == 'e' ||s[x] == 'i' ||s[x] == 'o' || s[x] == 'u' {
			tmpValue += 1
		}
		maxValue = max(maxValue , tmpValue)

		x ++
		y ++
	}
	return maxValue
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
