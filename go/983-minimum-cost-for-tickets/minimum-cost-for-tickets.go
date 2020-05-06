package main

func main() {
	a := []int {1,4,6,7,8,20}
	b := []int {2,7,15}
	mincostTickets(a , b)
}

func mincostTickets(days []int, costs []int) int {
	dp := make([]int , days[len(days) - 1] + 1)
	index := 0

	for i := 1 ; i < len(dp); i ++ {
		if i != days[index] {
			dp[i] = dp[i - 1]
		} else {
			dp[i] = min(dp[max(0, i - 1)] + costs[0], dp[max(0, i - 7)] + costs[1], dp[max(0, i - 30)] + costs[2])
			index ++
		}
	}
	return dp[len(dp) - 1]
}

func min(a, b, c int) int {
	min := a
	if b < min{
		min = b
	}
	if c < min {
		min = c
	}
	return min
}

func max(a ,b int) int  {
	if a < b {
		return b
	}
	return a
}