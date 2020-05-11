package leetcode


func MyPow(x float64, n int) float64 {
	var result float64
	result = recursion(x , n)
	if n < 0 {
		result = 1 / result
	}
	return result
}

func recursion(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := recursion(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}
