package leetcode

func pathWithObstacles(obstacleGrid [][]int) [][]int {
	r, c := len(obstacleGrid), len(obstacleGrid[0])
	visited := make([][]int, r)
	for i:=0; i<r; i++ {
		visited[i] = make([]int, c)
	}


	res := make([][]int, 0)

	var helper func(path [][]int, i, j int) bool
	helper = func(path [][]int, i, j int) bool {
		if visited[i][j] == 1 || obstacleGrid[i][j] == 1 {
			return false
		}
		if i==r-1 && j==c-1 {
			res = append(path, []int{i, j})
			return true
		}
		visited[i][j] = 1
		path = append(path, []int{i, j})
		if i<r-1 && helper(path, i+1, j) {
			return true
		}
		if j<c-1 && helper(path, i, j+1) {
			return true
		}
		path = path[:len(path)-1]
		return false
	}

	helper([][]int(nil), 0, 0)
	return res
}
