package leetcode


func CanMeasureWater(x int, y int, z int) bool {
	return z==0 || (x+y >= z && z % gcd(x, y) == 0)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a % b)
	}
}