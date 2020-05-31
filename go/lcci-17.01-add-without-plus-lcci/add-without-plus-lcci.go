package leetcode

func add(a int, b int) int {
	tmp := 0
	for b != 0 {
		tmp = a ^ b
		b = (a & b) << 1
		a = tmp
	}
	return a
}
