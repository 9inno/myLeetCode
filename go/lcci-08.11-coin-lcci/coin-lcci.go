package leetcode



func WaysToChange(n int) int {

	coin := []int{1, 5, 10, 25}
	dp := make([]int, n + 1)
	dp[0] = 1

	for i := 0 ; i < len(coin); i ++ {
		for j := 1; j <= n ; j ++ {
			if j - coin[i] >= 0 {
				dp[j] += dp[j - coin[i]]
			}
		}
	}
	return dp[n] % 1000000007

}