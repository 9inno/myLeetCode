package leetcode


func longestPalindrome(s string) string {
	length := len(s)
	result := ""
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, length)
	}
	for l := 0; l < length; l++ {
		for i := 0; i + l < length; i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = 1
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] > 0 && l + 1 > len(result) {
				result = s[i:i+l+1]
			}
		}
	}
	return result
}
