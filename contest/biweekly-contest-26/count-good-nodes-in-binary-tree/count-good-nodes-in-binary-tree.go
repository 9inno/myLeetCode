package leetcode

import "math"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	result := 0
	recursion(root, math.MinInt64, &result)
	return result
}

func recursion(root *TreeNode, maxValue int, result *int) {
	if root == nil {
		return
	}
	if root.Val >= maxValue {
		*result ++
	}
	recursion(root.Left, max(maxValue, root.Val), result)
	recursion(root.Right, max(maxValue, root.Val), result)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}