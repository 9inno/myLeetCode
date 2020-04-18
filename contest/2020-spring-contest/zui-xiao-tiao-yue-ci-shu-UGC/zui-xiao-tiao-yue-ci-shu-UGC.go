package main

import "fmt"

func main() {
	jump := []int{2, 7, 3, 1, 1, 1 , 1, 5, 1, 1, 1}
	fmt.Print(minJump(jump))
}

func minJump(jump []int) int {
	n := len(jump)
	visited := make([]bool, n)
	left := 0
	result := 0
	for q := []int{0}; len(q) > 0; {
		result++
		tmp := q
		q = []int{}
		for _ , current := range tmp {
			visited[current] = true
			right := current + jump[current]
			if right >= n {
				return result
			}
			if !visited[right] {
				q = append(q, right)
			}
			for ; left < current; left++ {
				if !visited[left] {
					q = append(q, left)
				}
			}
		}
	}
	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

