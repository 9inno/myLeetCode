package leetcode

import "sort"

func Merge(intervals [][]int) [][]int {

	sort.Sort(IntSlice(intervals))

	var tmp []int
	var result [][]int
	for _ , value := range intervals {
		if len(tmp) == 0 {
			tmp = value
			continue
		}
		if value[0] <= tmp[1] {
			tmp[1] = max(value[1], tmp[1])
		} else {
			result = append(result , tmp)
			tmp = value
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
type IntSlice [][]int

func (s IntSlice) Less(i, j int) bool { return s[i][0] < s[j][0]}
func (s IntSlice) Swap(i, j int) {s[i], s[j] = s[j], s[i]}
func (s IntSlice) Len() int { return len(s) }