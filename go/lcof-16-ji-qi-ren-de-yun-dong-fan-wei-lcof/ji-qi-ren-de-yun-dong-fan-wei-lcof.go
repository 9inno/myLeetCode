package leetcode

import (
	"strconv"
)

func MovingCount(m int, n int, k int) int {

	result := 0
	type void struct {}
	var visited void
	set := make(map[string]void)
	var queue []struct{ a, b int }
	queue = append(queue, struct{ a, b int }{a: 0, b: 0})

	for len(queue) != 0 {
		var tmpQueue []struct{ a, b int}
		for _ , element := range queue {
			key := strconv.Itoa(element.a) + "," + strconv.Itoa(element.b)
			if _, ok := set[key]; ok {
				continue
			}

			set[key] = visited
			sum := 0
			sum += getSum(element.a)
			sum += getSum(element.b)
			if sum <= k {
				result ++
				if element.a + 1 < m {
					tmpQueue = append(tmpQueue, struct{ a, b int }{a: element.a + 1, b: element.b})
				}
				if element.b + 1 < n {
					tmpQueue = append(tmpQueue, struct{ a, b int }{a: element.a, b: element.b + 1})
				}
			}

		}
		queue = tmpQueue
	}

	return result
}

func MovingCount1(m int, n int, k int) int {
	vis := make([][]bool, m + 1)
	for i := 0; i < len(vis); i++ {
		vis[i] = make([]bool, n + 1)
	}
	return dfs(0, 0, m, n, k, vis)
}

func dfs(x, y, m, n, k int, vis [][]bool) int {
	if getSum(x) + getSum(y) > k {
		return 0
	}
	if x >= m || y >= n || x < 0 || y < 0 {
		return 0
	}
	if vis[x][y] {
		return 0
	}
	vis[x][y] = true
	sum := 1
	sum += dfs(x - 1, y, m, n, k, vis)
	sum += dfs(x + 1, y, m, n, k, vis)
	sum += dfs(x, y - 1, m, n, k, vis)
	sum += dfs(x, y + 1, m, n, k, vis)
	return sum
}

func getSum(value int) int {
	sum := 0
	for value != 0 {
		sum += value % 10
		value /= 10
	}
	return sum
}