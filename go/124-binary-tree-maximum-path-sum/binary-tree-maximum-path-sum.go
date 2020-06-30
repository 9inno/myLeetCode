package leetcode

import "math"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	result := math.MinInt32
	recursion(root, &result)
	return result
}

func recursion(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}
	left := max(recursion(root.Left, result), 0)
	right := max(recursion(root.Right, result), 0)
	*result = max(*result, left + right + root.Val)
	return root.Val + max(left , right)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
