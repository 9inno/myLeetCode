package	leetcode

func MaxDistance(grid [][]int) int {

	res := -1
	hasLevel := true;
	level := 1
	for hasLevel {

		hasLevel = false
		for i := 0; i< len(grid); i++ {
			for j := 0; j < len(grid); j++ {
				if i < len(grid)-1 && j<len(grid)-1 &&
					( grid[i][j] != grid[i+1][j] || grid[i][j] != grid[i][j+1]) {
					hasLevel = true
				}else if i == len(grid) - 1&& j<len(grid)-1 &&
					(grid[i][j] != grid[i][j+1]){
					hasLevel = true
				}else if i < len(grid)-1 && j == len(grid) -1 &&
					(grid[i+1][j] != grid[i][j]){
					hasLevel = true
				}
				if grid[i][j] == level && hasLevel {
					dfs(grid, i, j, level)

				}
			}
		}
		if hasLevel {
			res = level
			level++
		}
	}

	return res
}

func dfs(grid [][]int, i, j, level int) {

	if i < 0 || i > len(grid)-1 || j < 0 || j > len(grid)-1 || grid[i][j] == level+1{
		return
	}
	if grid[i][j] == 0 {
		grid[i][j] = level+1
		return
	}
	grid[i][j] = level+1
	dfs(grid, i - 1, j, level)
	dfs(grid, i + 1, j, level)
	dfs(grid, i , j - 1, level)
	dfs(grid, i , j + 1, level)
}