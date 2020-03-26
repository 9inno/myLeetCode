package leetcode


func SurfaceArea(grid [][]int) int {
	result := 0
	lenY := len(grid)
	if lenY == 0 {
		return result
	}
	lenX := len(grid[0])
	for i := 0; i < lenY; i ++ {
		for j := 0; j < lenX ; j ++ {
			if grid[i][j] == 0 {
				continue
			}
			result += grid[i][j] * 4 + 2
			if j+1 < lenX {
				result -= min(grid[i][j+1], grid[i][j]) * 2
			}
			if i+1 < lenY {
				result -= min(grid[i+1][j], grid[i][j]) * 2
			}

		}
	}
	return result
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}