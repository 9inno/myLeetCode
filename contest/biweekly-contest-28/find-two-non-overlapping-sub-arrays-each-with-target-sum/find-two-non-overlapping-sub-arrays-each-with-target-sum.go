package main

import (
	"sort"
)

func main() {
	//a := []int{78,18,1,94,1,1,1,29,58,3,4,1,2,56,17,19,4,1,63,2,16,11,1,1,2,1,25,62,10,69,12,7,1,6,2,92,4,1,61,7,26,1,1,1,67,26,2,2,70,25,2,68,13,4,11,1,34,14,7,37,4,1,12,51,25,2,4,3,56,21,7,8,5,93,1,1,2,55,14,25,1,1,1,89,6,1,1,24,22,50,1,28,9,51,9,88,1,7,1,30,32,18,12,3,2,18,10,4,11,43,6,5,93,2,2,68,18,11,47,33,17,27,56,13,1,2,29,1,17,1,10,15,18,3,1,86,7,4,16,45,3,29,2,1,1,31,19,18,16,12,1,56,4,35,1,1,36,59,1,1,16,58,18,4,1,43,31,15,6,1,1,6,49,27,12,1,2,80,14,2,1,21,32,18,15,11,59,10,1,14,3,3,7,15,4,55,4,1,12,4,1,1,53,37,2,5,72,3,6,10,3,3,83,8,1,5}
	a := []int{47,17,4,8,8,2,1,1,8,35,30,1,11,1,1,1,44,1,3,27,2,8}
	minSumOfLengths(a , 88)
}

func minSumOfLengths(arr []int, target int) int {
	sort.Ints(arr)
	var ans [][]int
	backtrack(arr, target, 0, []int{}, &ans)
	sort.Slice(ans, func(i, j int) bool {

	})
	return 1
}
//
//func min(a, b int) int {
//	if a > b {
//		return a
//	}
//	return a
//}



func backtrack(candidates []int, target, index int, path []int, ans *[][]int) {
	// 退出条件
	if target == 0 {
		*ans = append(*ans, path)
		return
	}
	n := len(candidates)
	for i := index; i < n; i++ {
		// 去除多余的重复项（调整当前层的选择进度即可，下层需要继续可选）
		rawI := i
		for i < n - 1 && candidates[i] == candidates[i+1] {
			i++
		}
		// 剪枝
		if candidates[i] > target {
			return
		}
		// 作出选择
		newPath := make([]int, len(path))
		copy(newPath, path)
		newPath = append(newPath, candidates[i])
		backtrack(candidates, target-candidates[i], rawI+1, newPath, ans)
		// 撤回选择

	}
}
