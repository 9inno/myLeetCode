package leetcode

func kthFactor(n int, k int) int {
	result := []int{}
	for i := 1; i <= n ;i ++ {
		if n % i == 0 {
			result = append(result, i)
		}
	}
	if k > len(result) {
		return  -1
	}
	return result[k - 1]
}
