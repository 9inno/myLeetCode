package leetcode

func CoinChange(coins []int, amount int) int {
	dp := make([]int , amount + 1)
	for i := 0; i <= amount; i++ { dp[i] = amount+1 }
	dp[0] = 0
	for i := 0; i <= amount; i = i + 1 {
		for j := 0; j < len(coins); j = j + 1 {
			if coins[j] <= i {
				if dp[i] > dp[i-coins[j]]+1 {
					dp[i] = dp[i-coins[j]]+1
				}
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}


