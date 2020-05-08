package leetcode


func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	length:= 0
	for i := 0; i < len(matrix); i ++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0 ; j < len(matrix[i]); j ++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				length = 1
			}
		}
	}

	for i := 1; i < len(dp); i ++ {
		for j := 1 ; j < len(dp[i]); j ++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(dp[i - 1][j], dp[i][j - 1], dp[i - 1][j - 1]) + 1
				if dp[i][j] > length {
					length =dp[i][j]
				}
			}

		}
	}
	
	return  length * length

}

func min(a, b, c int) int {
	minValue := a
	if b < minValue {
		minValue = b
	}
	if c < minValue {
		minValue = c
	}
	return minValue
}