package leetcode

func IsHappy(n int) bool {
	result := 0
	for n != 0 {
		tmp := n % 10
		result += tmp * tmp
		n /= 10
	}
	if result > 4 {
		return IsHappy(result)
	}
	return result == 1
}
