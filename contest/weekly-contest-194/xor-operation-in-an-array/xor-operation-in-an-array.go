package leetcode

func xorOperation(n int, start int) int {
	result :=0
	for i := 0 ; i < n ; i ++ {
		result ^= i + 2 * start
	}
	return result
}
