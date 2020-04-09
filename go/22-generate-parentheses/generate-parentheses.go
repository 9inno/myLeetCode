package leetcode



func GenerateParenthesis(n int) []string {

	var result []string
	reversion(&result,n,n, "")
	return result
}

func reversion(result *[]string, left, right int, current string) {

	if left > right {
		return
	}

	if left == 0 && right == 0 {
		*result = append(*result , current)
	}

	if left > 0 {
		reversion(result, left - 1, right, current + "(")
	}

	if right > 0 {
		reversion(result, left, right - 1, current + ")")
	}


}