package leetcode

func multiply(A int, B int) int {
	result := 0
	if A > B {
		result = recursion(A, 1, B)
	} else {
		result = recursion(B, 1, A)
	}
	return result
}

func recursion(A, start, end int) int {
	if end == 0 {
		return 0
	}
	if start >= end {
		return A
	}
	mid := start + (end - start) / 2
	left := recursion(A, start, mid)
	right := recursion(A, mid + 1, end)

	return left + right
}