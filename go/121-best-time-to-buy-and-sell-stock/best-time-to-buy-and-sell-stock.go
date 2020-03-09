package leetcode


func MaxProfit(prices []int) int {
	result := 0
	if len(prices) == 0 {
		return result
	}
	buy := prices[0]
	for i := 0; i < len(prices); i = i + 1 {
		if buy > prices[i] {
			buy = prices[i]
		}else {
			if result <  prices[i] - buy  {
				result = prices[i] - buy
			}
		}
	}
	return result
}