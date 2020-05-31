package leetcode

import "sort"

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	maxH := horizontalCuts[0]
	maxV := verticalCuts[0]

	for i := 1 ; i < len(horizontalCuts); i ++ {
		maxH = max(maxH, horizontalCuts[i] -horizontalCuts[i - 1])
	}
	for i := 1 ; i < len(verticalCuts); i ++ {
		maxV = max(maxV, verticalCuts[i] - verticalCuts[ i - 1])
	}
	maxH = max(maxH, h - horizontalCuts[len(horizontalCuts) - 1])
	maxV = max(maxV, w - verticalCuts[len(verticalCuts) - 1])

	return (maxH* maxV) % 1000000007

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}