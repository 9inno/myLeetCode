package leetocde

import "math"

func Maximum(a int, b int) int {
	return (a + b + int(math.Abs(float64(a - b)))) >> 1
}