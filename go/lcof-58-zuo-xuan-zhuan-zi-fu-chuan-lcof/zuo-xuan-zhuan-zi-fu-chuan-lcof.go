package leetcode


func ReverseLeftWords(s string, n int) string {
	if len(s) < n {
		return s
	}
	return s[n:] + s[0:n]
}

