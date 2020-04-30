package leetcode

import "sort"

func Insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var tmp []int
	var result [][]int
	for i := 0 ; i < len(intervals) ; i ++ {
		if len(tmp) == 0 {
			tmp = intervals[i]
			continue
		}
		if intervals[i][0] <= tmp[1] {
			tmp[1] = max(intervals[i][1], tmp[1])
		} else {
			result = append(result , tmp)
			tmp = intervals[i]
		}
	}
	if tmp != nil {
		result = append(result , tmp)
	}
	return result
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}