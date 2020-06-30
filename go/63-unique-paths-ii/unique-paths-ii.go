package leetcode


func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}
	result := make([][]int, len(obstacleGrid))
	jFlag := true
	for i := 0; i < len(obstacleGrid); i ++ {
		result[i] = make([]int, len(obstacleGrid[0]))
		for j := 0; j < len(obstacleGrid[0]); j ++ {
			if i == 0 {
				if obstacleGrid[i][j] == 1 {
					break
				}
				result[i][j] = 1
				continue
			}

			if j == 0 {
				if obstacleGrid[i][j] == 1 {
					jFlag = false
				}
				if jFlag {
					result[i][j] = 1
				} else {
					result[i][j] = 0
				}
				continue
			}
			if obstacleGrid[i][j] == 1 {
				result[i][j] = 0
				continue
			}
			result[i][j] = result[i - 1][j] + result[i][j - 1]

		}
	}
	return result[len(result) - 1][len(result[len(result) - 1]) - 1]
}
