package leetcode

func WaysToStep(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return n
	}
	if n == 3 {
		return 4
	}

	a := 1
	b := 2
	c := 4
	for i := 4 ; i < n + 1; i ++ {
		tmp := (a + b + c) % 1000000007
		a = b
		b = c
		c = tmp
	}
	return c
}
