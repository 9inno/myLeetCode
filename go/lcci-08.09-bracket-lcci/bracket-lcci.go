package leetcode

func generateParenthesis(n int) []string {
	result := []string{}
	generate("", n, n , &result)
	return result
}

func generate(current string, left, right int, result *[]string) {

	if left > right {
		return
	}
	if left == 0 && right == 0 {
		*result = append(*result, current)
		return
	}
	if left > 0 {
		generate(current+ "(", left - 1 , right , result)
	}
	if right > 0 {
		generate(current+ ")", left , right - 1 , result)
	}
}
